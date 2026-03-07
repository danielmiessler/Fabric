package pipeline

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/strategy"
)

var envReferencePattern = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)\}|\$([A-Za-z_][A-Za-z0-9_]*)`)

type PreflightOptions struct {
	Registry *core.PluginRegistry
}

func Preflight(ctx context.Context, p *Pipeline, opts PreflightOptions) error {
	if err := Validate(p); err != nil {
		return err
	}
	if ctx != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	for i := range p.Stages {
		stage := &p.Stages[i]

		var err error
		stage.Pattern, err = expandEnvReferences(stage.Pattern)
		if err != nil {
			return fmt.Errorf("pipeline %q stage %q pattern: %w", p.Name, stage.ID, err)
		}
		stage.Context, err = expandEnvReferences(stage.Context)
		if err != nil {
			return fmt.Errorf("pipeline %q stage %q context: %w", p.Name, stage.ID, err)
		}
		stage.Strategy, err = expandEnvReferences(stage.Strategy)
		if err != nil {
			return fmt.Errorf("pipeline %q stage %q strategy: %w", p.Name, stage.ID, err)
		}
		for key, value := range stage.Variables {
			stage.Variables[key], err = expandEnvReferences(value)
			if err != nil {
				return fmt.Errorf("pipeline %q stage %q variable %q: %w", p.Name, stage.ID, key, err)
			}
		}
		for artifactIndex := range stage.Artifacts {
			stage.Artifacts[artifactIndex].Path, err = expandEnvReferences(stage.Artifacts[artifactIndex].Path)
			if err != nil {
				return fmt.Errorf("pipeline %q stage %q artifact %q: %w", p.Name, stage.ID, stage.Artifacts[artifactIndex].Name, err)
			}
		}

		switch stage.Executor {
		case ExecutorBuiltin:
			if !isSupportedBuiltin(stage.Builtin.Name) {
				return fmt.Errorf("pipeline %q stage %q builtin %q is not supported", p.Name, stage.ID, stage.Builtin.Name)
			}
		case ExecutorCommand:
			if err := preflightCommandStage(p, stage); err != nil {
				return err
			}
		case ExecutorFabricPattern:
			if err := preflightPatternStage(p, stage, opts.Registry); err != nil {
				return err
			}
		}
	}

	return nil
}

func preflightCommandStage(p *Pipeline, stage *Stage) error {
	var err error
	stage.Command.Program, err = expandEnvReferences(stage.Command.Program)
	if err != nil {
		return fmt.Errorf("pipeline %q stage %q command.program: %w", p.Name, stage.ID, err)
	}
	stage.Command.Cwd, err = expandEnvReferences(stage.Command.Cwd)
	if err != nil {
		return fmt.Errorf("pipeline %q stage %q command.cwd: %w", p.Name, stage.ID, err)
	}
	for i := range stage.Command.Args {
		stage.Command.Args[i], err = expandEnvReferences(stage.Command.Args[i])
		if err != nil {
			return fmt.Errorf("pipeline %q stage %q command.args[%d]: %w", p.Name, stage.ID, i, err)
		}
	}
	for key, value := range stage.Command.Env {
		stage.Command.Env[key], err = expandEnvReferences(value)
		if err != nil {
			return fmt.Errorf("pipeline %q stage %q command.env.%s: %w", p.Name, stage.ID, key, err)
		}
	}

	if !containsRuntimePlaceholder(stage.Command.Program) {
		if resolvedProgram, err := resolveCommandProgram(p, stage.Command.Program); err != nil {
			return fmt.Errorf("pipeline %q stage %q command.program: %w", p.Name, stage.ID, err)
		} else {
			stage.Command.Program = resolvedProgram
		}
	}
	if stage.Command.Cwd != "" && !containsRuntimePlaceholder(stage.Command.Cwd) {
		resolvedCwd := resolvePipelinePath(p, stage.Command.Cwd)
		info, err := os.Stat(resolvedCwd)
		if err != nil {
			return fmt.Errorf("pipeline %q stage %q command.cwd: %w", p.Name, stage.ID, err)
		}
		if !info.IsDir() {
			return fmt.Errorf("pipeline %q stage %q command.cwd %q is not a directory", p.Name, stage.ID, resolvedCwd)
		}
		stage.Command.Cwd = resolvedCwd
	}

	return nil
}

func preflightPatternStage(p *Pipeline, stage *Stage, registry *core.PluginRegistry) error {
	if registry == nil {
		return fmt.Errorf("pipeline %q stage %q cannot validate patterns without plugin registry", p.Name, stage.ID)
	}
	if _, err := registry.Db.Patterns.Get(stage.Pattern); err != nil {
		return fmt.Errorf("pipeline %q stage %q pattern %q: %w", p.Name, stage.ID, stage.Pattern, err)
	}
	if stage.Context != "" {
		if _, err := registry.Db.Contexts.Get(stage.Context); err != nil {
			return fmt.Errorf("pipeline %q stage %q context %q: %w", p.Name, stage.ID, stage.Context, err)
		}
	}
	if stage.Strategy != "" {
		if _, err := strategy.LoadStrategy(stage.Strategy); err != nil {
			return fmt.Errorf("pipeline %q stage %q strategy %q: %w", p.Name, stage.ID, stage.Strategy, err)
		}
	}
	return nil
}

func expandEnvReferences(value string) (string, error) {
	if value == "" {
		return value, nil
	}
	matches := envReferencePattern.FindAllStringSubmatch(value, -1)
	for _, match := range matches {
		name := match[1]
		if name == "" {
			name = match[2]
		}
		if _, exists := os.LookupEnv(name); !exists {
			return "", fmt.Errorf("missing environment variable %q", name)
		}
	}
	return os.ExpandEnv(value), nil
}

func resolveCommandProgram(p *Pipeline, program string) (string, error) {
	if program == "" {
		return "", nil
	}
	if filepath.IsAbs(program) {
		if _, err := os.Stat(program); err != nil {
			return "", err
		}
		return program, nil
	}
	if looksLikeRelativePath(program) {
		resolved := resolvePipelinePath(p, program)
		if _, err := os.Stat(resolved); err != nil {
			return "", err
		}
		return resolved, nil
	}
	resolved, err := exec.LookPath(program)
	if err != nil {
		return "", err
	}
	return resolved, nil
}

func resolvePipelinePath(p *Pipeline, value string) string {
	if value == "" || filepath.IsAbs(value) || p.FilePath == "" {
		return value
	}
	return filepath.Clean(filepath.Join(filepath.Dir(p.FilePath), value))
}

func looksLikeRelativePath(value string) bool {
	return strings.HasPrefix(value, ".") || strings.ContainsRune(value, filepath.Separator)
}

func containsRuntimePlaceholder(value string) bool {
	return runtimePlaceholderPattern.MatchString(value)
}

func isSupportedBuiltin(name string) bool {
	switch name {
	case "passthrough", "noop", "source_capture", "validate_declared_outputs", "write_publish_manifest":
		return true
	default:
		return false
	}
}
