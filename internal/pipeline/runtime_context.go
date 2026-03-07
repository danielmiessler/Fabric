package pipeline

import (
	"fmt"
	"regexp"
	"strings"
)

var runtimePlaceholderPattern = regexp.MustCompile(`\{\{([^{}]+)\}\}`)

type StageRuntimeContext struct {
	Pipeline       *Pipeline
	Stage          Stage
	Source         RunSource
	InputPayload   string
	InvocationDir  string
	RunDir         string
	RunID          string
	StageArtifacts map[string]map[string]string
	StagePayloads  map[string]string
	Manifest       *RunManifest
	FinalOutput    string
}

func interpolateRuntimeValue(value string, runtimeCtx StageRuntimeContext) (string, error) {
	if value == "" {
		return value, nil
	}

	var interpolationErr error
	result := runtimePlaceholderPattern.ReplaceAllStringFunc(value, func(token string) string {
		if interpolationErr != nil {
			return token
		}

		key := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(token, "{{"), "}}"))
		resolved, err := resolveRuntimePlaceholder(key, runtimeCtx)
		if err != nil {
			interpolationErr = err
			return token
		}
		return resolved
	})
	if interpolationErr != nil {
		return "", interpolationErr
	}
	return result, nil
}

func resolveRuntimePlaceholder(key string, runtimeCtx StageRuntimeContext) (string, error) {
	switch key {
	case "source":
		if runtimeCtx.Source.Reference != "" {
			return runtimeCtx.Source.Reference, nil
		}
		return runtimeCtx.Source.Payload, nil
	case "source_reference":
		return runtimeCtx.Source.Reference, nil
	case "source_payload", "input", "previous":
		return runtimeCtx.InputPayload, nil
	case "run_dir":
		return runtimeCtx.RunDir, nil
	case "run_id":
		return runtimeCtx.RunID, nil
	case "invocation_dir":
		return runtimeCtx.InvocationDir, nil
	case "stage_id":
		return runtimeCtx.Stage.ID, nil
	}

	if strings.HasPrefix(key, "artifact:") {
		parts := strings.Split(key, ":")
		if len(parts) != 3 {
			return "", fmt.Errorf("invalid runtime placeholder %q", key)
		}
		stageArtifacts := runtimeCtx.StageArtifacts[parts[1]]
		if stageArtifacts == nil {
			return "", fmt.Errorf("runtime placeholder %q references unknown stage", key)
		}
		path := stageArtifacts[parts[2]]
		if path == "" {
			return "", fmt.Errorf("runtime placeholder %q references unknown artifact", key)
		}
		return path, nil
	}

	return "", fmt.Errorf("unknown runtime placeholder %q", key)
}

func effectiveStageRole(stage Stage) StageRole {
	if stage.Role != StageRoleDefault {
		return stage.Role
	}
	if stage.Executor == ExecutorBuiltin && stage.Builtin != nil {
		switch stage.Builtin.Name {
		case "validate_declared_outputs":
			return StageRoleValidate
		case "write_publish_manifest":
			return StageRolePublish
		}
	}
	return StageRoleDefault
}

func findLastValidateStageIndex(stages []Stage) int {
	last := -1
	for i, stage := range stages {
		if effectiveStageRole(stage) == StageRoleValidate {
			last = i
		}
	}
	return last
}
