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

func (r *Runner) executeCommandStage(ctx context.Context, stage Stage, runtimeCtx StageRuntimeContext) (*StageExecutionResult, error) {
	commandCtx := ctx
	cancel := func() {}
	if stage.Command.Timeout > 0 {
		commandCtx, cancel = context.WithTimeout(ctx, time.Duration(stage.Command.Timeout)*time.Second)
	}
	defer cancel()

	commandConfig, err := resolveCommandConfig(stage, runtimeCtx)
	if err != nil {
		return nil, err
	}
	cmd := exec.CommandContext(commandCtx, commandConfig.Program, commandConfig.Args...)
	cmd.Dir = runtimeCtx.InvocationDir
	if commandConfig.Cwd != "" {
		cmd.Dir = commandConfig.Cwd
	}

	cmd.Env = buildCommandEnv(stage, runtimeCtx.RunDir, runtimeCtx.Source, runtimeCtx.InvocationDir, commandConfig.Env)
	cmd.Stdin = strings.NewReader(runtimeCtx.InputPayload)

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

func resolveCommandConfig(stage Stage, runtimeCtx StageRuntimeContext) (*CommandConfig, error) {
	resolved := &CommandConfig{
		Args:    make([]string, len(stage.Command.Args)),
		Env:     make(map[string]string, len(stage.Command.Env)),
		Timeout: stage.Command.Timeout,
	}

	var err error
	resolved.Program, err = interpolateRuntimeValue(stage.Command.Program, runtimeCtx)
	if err != nil {
		return nil, fmt.Errorf("command stage %q program: %w", stage.ID, err)
	}
	resolved.Cwd, err = interpolateRuntimeValue(stage.Command.Cwd, runtimeCtx)
	if err != nil {
		return nil, fmt.Errorf("command stage %q cwd: %w", stage.ID, err)
	}
	for i, arg := range stage.Command.Args {
		resolved.Args[i], err = interpolateRuntimeValue(arg, runtimeCtx)
		if err != nil {
			return nil, fmt.Errorf("command stage %q arg[%d]: %w", stage.ID, i, err)
		}
	}
	for key, value := range stage.Command.Env {
		resolved.Env[key], err = interpolateRuntimeValue(value, runtimeCtx)
		if err != nil {
			return nil, fmt.Errorf("command stage %q env.%s: %w", stage.ID, key, err)
		}
	}

	return resolved, nil
}

func buildCommandEnv(stage Stage, runDir string, source RunSource, invocationDir string, stageEnv map[string]string) []string {
	env := make(map[string]string, len(stageEnv)+7)
	for _, item := range os.Environ() {
		key, value, found := strings.Cut(item, "=")
		if !found {
			continue
		}
		env[key] = value
	}
	env["FABRIC_PIPELINE_RUN_DIR"] = runDir
	env["FABRIC_PIPELINE_STAGE_ID"] = stage.ID
	env["FABRIC_PIPELINE_SOURCE_MODE"] = string(source.Mode)
	env["FABRIC_PIPELINE_SOURCE_REFERENCE"] = source.Reference
	env["FABRIC_PIPELINE_INVOCATION_DIR"] = invocationDir
	for _, artifact := range stage.Artifacts {
		env[artifactEnvKey(artifact.Name)] = filepath.Join(runDir, artifact.Path)
	}
	for key, value := range stageEnv {
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

func (r *Runner) executePatternStage(ctx context.Context, stage Stage, runtimeCtx StageRuntimeContext) (*StageExecutionResult, error) {
	if r.registry == nil {
		return nil, fmt.Errorf("pattern stage %q cannot run without plugin registry", stage.ID)
	}

	chatOptions := &domain.ChatOptions{
		Quiet: true,
	}
	if stage.Stream {
		chatOptions.UpdateChan = make(chan domain.StreamUpdate)
	}

	strategyName, err := interpolateRuntimeValue(stage.Strategy, runtimeCtx)
	if err != nil {
		return nil, fmt.Errorf("pattern stage %q strategy: %w", stage.ID, err)
	}

	chatter, err := r.registry.GetChatter("", 0, "", strategyName, stage.Stream, false)
	if err != nil {
		return nil, err
	}

	contextName, err := interpolateRuntimeValue(stage.Context, runtimeCtx)
	if err != nil {
		return nil, fmt.Errorf("pattern stage %q context: %w", stage.ID, err)
	}
	patternName, err := interpolateRuntimeValue(stage.Pattern, runtimeCtx)
	if err != nil {
		return nil, fmt.Errorf("pattern stage %q pattern: %w", stage.ID, err)
	}
	variables := make(map[string]string, len(stage.Variables))
	for key, value := range stage.Variables {
		variables[key], err = interpolateRuntimeValue(value, runtimeCtx)
		if err != nil {
			return nil, fmt.Errorf("pattern stage %q variable %q: %w", stage.ID, key, err)
		}
	}

	req := &domain.ChatRequest{
		ContextName:      contextName,
		PatternName:      patternName,
		PatternVariables: variables,
		StrategyName:     strategyName,
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: strings.TrimSpace(runtimeCtx.InputPayload),
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
