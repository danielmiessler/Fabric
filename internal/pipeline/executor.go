package pipeline

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

func (r *Runner) executeCommandStage(ctx context.Context, p *Pipeline, stage Stage, inputPayload string, runDir string, invocationDir string) (*StageExecutionResult, error) {
	commandCtx := ctx
	cancel := func() {}
	if stage.Command.Timeout > 0 {
		commandCtx, cancel = context.WithTimeout(ctx, time.Duration(stage.Command.Timeout)*time.Second)
	}
	defer cancel()

	cmd := exec.CommandContext(commandCtx, stage.Command.Program, stage.Command.Args...)
	cmd.Dir = invocationDir
	if stage.Command.Cwd != "" {
		cmd.Dir = stage.Command.Cwd
	}

	cmd.Env = buildCommandEnv(stage, runDir)
	cmd.Stdin = strings.NewReader(inputPayload)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = io.MultiWriter(r.Stderr)

	if err := cmd.Run(); err != nil {
		if commandCtx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("command stage %q timed out after %d seconds", stage.ID, stage.Command.Timeout)
		}
		return nil, fmt.Errorf("command stage %q failed: %w", stage.ID, err)
	}

	return &StageExecutionResult{Stdout: stdout.String()}, nil
}

func buildCommandEnv(stage Stage, runDir string) []string {
	env := make(map[string]string, len(stage.Command.Env)+4)
	for _, item := range os.Environ() {
		key, value, found := strings.Cut(item, "=")
		if !found {
			continue
		}
		env[key] = value
	}
	env["FABRIC_PIPELINE_RUN_DIR"] = runDir
	env["FABRIC_PIPELINE_STAGE_ID"] = stage.ID
	for _, artifact := range stage.Artifacts {
		env[artifactEnvKey(artifact.Name)] = filepath.Join(runDir, artifact.Path)
	}
	for key, value := range stage.Command.Env {
		env[key] = value
	}

	keys := make([]string, 0, len(env))
	for key := range env {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+"="+env[key])
	}
	return result
}

func artifactEnvKey(name string) string {
	replacer := strings.NewReplacer("-", "_", ".", "_", " ", "_", "/", "_")
	return "FABRIC_PIPELINE_ARTIFACT_" + strings.ToUpper(replacer.Replace(name))
}

func (r *Runner) executePatternStage(ctx context.Context, stage Stage, inputPayload string) (*StageExecutionResult, error) {
	if r.registry == nil {
		return nil, fmt.Errorf("pattern stage %q cannot run without plugin registry", stage.ID)
	}

	chatOptions := &domain.ChatOptions{
		Quiet: true,
	}
	if stage.Stream {
		chatOptions.UpdateChan = make(chan domain.StreamUpdate)
	}

	chatter, err := r.registry.GetChatter("", 0, "", stage.Strategy, stage.Stream, false)
	if err != nil {
		return nil, err
	}

	req := &domain.ChatRequest{
		ContextName:      stage.Context,
		PatternName:      stage.Pattern,
		PatternVariables: stage.Variables,
		StrategyName:     stage.Strategy,
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: strings.TrimSpace(inputPayload),
		},
	}

	if stage.Stream {
		errCh := make(chan error, 1)
		go func() {
			for update := range chatOptions.UpdateChan {
				switch update.Type {
				case domain.StreamTypeContent:
					fmt.Fprint(r.Stderr, update.Content)
				case domain.StreamTypeError:
					errCh <- errors.New(update.Content)
				}
			}
			close(errCh)
		}()

		session, sendErr := chatter.Send(req, chatOptions)
		close(chatOptions.UpdateChan)
		if sendErr != nil {
			for range errCh {
			}
			return nil, sendErr
		}
		for streamErr := range errCh {
			if streamErr != nil {
				return nil, streamErr
			}
		}
		return &StageExecutionResult{Stdout: session.GetLastMessage().Content}, nil
	}

	session, err := chatter.Send(req, chatOptions)
	if err != nil {
		return nil, err
	}
	return &StageExecutionResult{Stdout: session.GetLastMessage().Content}, nil
}
