package pipeline

import (
	"fmt"
	"path/filepath"
)

func Validate(p *Pipeline) error {
	if p == nil {
		return fmt.Errorf("pipeline is nil")
	}
	if p.Version <= 0 {
		return fmt.Errorf("pipeline %q must declare version >= 1", p.FilePath)
	}
	if p.Name == "" {
		return fmt.Errorf("pipeline %q must declare name", p.FilePath)
	}
	if p.FileStem != "" && p.Name != p.FileStem {
		return fmt.Errorf("pipeline name %q must match filename stem %q", p.Name, p.FileStem)
	}
	if len(p.Stages) == 0 {
		return fmt.Errorf("pipeline %q must declare at least one stage", p.Name)
	}

	for _, mode := range p.Accepts {
		switch mode {
		case SourceModeStdin, SourceModeSource, SourceModeScrapeURL:
		default:
			return fmt.Errorf("pipeline %q declares unsupported source mode %q", p.Name, mode)
		}
	}

	stageByID := make(map[string]*Stage, len(p.Stages))
	finalOutputCount := 0
	for i := range p.Stages {
		stage := &p.Stages[i]

		if stage.ID == "" {
			return fmt.Errorf("pipeline %q has stage without id", p.Name)
		}
		if _, exists := stageByID[stage.ID]; exists {
			return fmt.Errorf("pipeline %q has duplicate stage id %q", p.Name, stage.ID)
		}
		stageByID[stage.ID] = stage

		switch stage.Role {
		case StageRoleDefault, StageRoleValidate, StageRolePublish:
		default:
			return fmt.Errorf("pipeline %q stage %q has unsupported role %q", p.Name, stage.ID, stage.Role)
		}

		switch stage.Executor {
		case ExecutorBuiltin:
			if stage.Builtin == nil || stage.Builtin.Name == "" {
				return fmt.Errorf("pipeline %q stage %q must declare builtin.name", p.Name, stage.ID)
			}
		case ExecutorCommand:
			if stage.Command == nil {
				return fmt.Errorf("pipeline %q stage %q must declare command", p.Name, stage.ID)
			}
			if stage.Command.Program == "" {
				return fmt.Errorf("pipeline %q stage %q must declare command.program", p.Name, stage.ID)
			}
			if stage.Command.Timeout < 0 {
				return fmt.Errorf("pipeline %q stage %q timeout must be >= 0", p.Name, stage.ID)
			}
		case ExecutorFabricPattern:
			if stage.Pattern == "" {
				return fmt.Errorf("pipeline %q stage %q must declare pattern", p.Name, stage.ID)
			}
		default:
			return fmt.Errorf("pipeline %q stage %q has unsupported executor %q", p.Name, stage.ID, stage.Executor)
		}

		artifactNames := make(map[string]struct{}, len(stage.Artifacts))
		for _, artifact := range stage.Artifacts {
			if artifact.Name == "" {
				return fmt.Errorf("pipeline %q stage %q declares artifact without name", p.Name, stage.ID)
			}
			if _, exists := artifactNames[artifact.Name]; exists {
				return fmt.Errorf("pipeline %q stage %q declares duplicate artifact name %q", p.Name, stage.ID, artifact.Name)
			}
			artifactNames[artifact.Name] = struct{}{}
			if artifact.Path == "" {
				return fmt.Errorf("pipeline %q stage %q artifact %q must declare path", p.Name, stage.ID, artifact.Name)
			}
			if filepath.IsAbs(artifact.Path) {
				return fmt.Errorf("pipeline %q stage %q artifact %q path must be relative", p.Name, stage.ID, artifact.Name)
			}
		}

		if stage.Input != nil {
			switch stage.Input.From {
			case StageInputSource, StageInputPrevious:
			case StageInputArtifact:
				if stage.Input.Stage == "" {
					return fmt.Errorf("pipeline %q stage %q input.from=artifact requires input.stage", p.Name, stage.ID)
				}
				if stage.Input.Artifact == "" {
					return fmt.Errorf("pipeline %q stage %q input.from=artifact requires input.artifact", p.Name, stage.ID)
				}
			default:
				return fmt.Errorf("pipeline %q stage %q has unsupported input.from %q", p.Name, stage.ID, stage.Input.From)
			}
		}

		if stage.PrimaryOutput != nil {
			switch stage.PrimaryOutput.From {
			case PrimaryOutputStdout:
			case PrimaryOutputArtifact:
				if stage.PrimaryOutput.Artifact == "" {
					return fmt.Errorf("pipeline %q stage %q primary_output.from=artifact requires artifact", p.Name, stage.ID)
				}
				if _, exists := artifactNames[stage.PrimaryOutput.Artifact]; !exists {
					return fmt.Errorf("pipeline %q stage %q primary_output artifact %q is not declared", p.Name, stage.ID, stage.PrimaryOutput.Artifact)
				}
			default:
				return fmt.Errorf("pipeline %q stage %q has unsupported primary_output.from %q", p.Name, stage.ID, stage.PrimaryOutput.From)
			}
		}

		if stage.FinalOutput {
			finalOutputCount++
			if stage.PrimaryOutput == nil {
				return fmt.Errorf("pipeline %q stage %q is final_output but has no primary_output", p.Name, stage.ID)
			}
		}
	}

	if finalOutputCount != 1 {
		return fmt.Errorf("pipeline %q must declare exactly one final_output stage", p.Name)
	}

	for i := range p.Stages {
		stage := &p.Stages[i]
		if stage.Input != nil && stage.Input.From == StageInputArtifact {
			targetStage, exists := stageByID[stage.Input.Stage]
			if !exists {
				return fmt.Errorf("pipeline %q stage %q references unknown stage %q", p.Name, stage.ID, stage.Input.Stage)
			}
			if !stageDeclaresArtifact(targetStage, stage.Input.Artifact) {
				return fmt.Errorf("pipeline %q stage %q references unknown artifact %q on stage %q", p.Name, stage.ID, stage.Input.Artifact, stage.Input.Stage)
			}
		}
	}

	return nil
}

func stageDeclaresArtifact(stage *Stage, name string) bool {
	for _, artifact := range stage.Artifacts {
		if artifact.Name == name {
			return true
		}
	}
	return false
}
