package pipeline

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateRejectsAdditionalInvalidPipelineContracts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		build    func() *Pipeline
		contains string
	}{
		{
			name: "nil pipeline",
			build: func() *Pipeline {
				return nil
			},
			contains: "pipeline is nil",
		},
		{
			name: "missing version",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Version = 0
				return p
			},
			contains: "must declare version >= 1",
		},
		{
			name: "missing name",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Name = ""
				return p
			},
			contains: "must declare name",
		},
		{
			name: "unsupported accepts mode",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Accepts = []SourceMode{SourceModeStdin, SourceMode("archive")}
				return p
			},
			contains: `declares unsupported source mode "archive"`,
		},
		{
			name: "no stages",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages = nil
				return p
			},
			contains: "must declare at least one stage",
		},
		{
			name: "stage missing id",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].ID = ""
				return p
			},
			contains: "stage without id",
		},
		{
			name: "unsupported role",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].Role = StageRole("shipit")
				return p
			},
			contains: "unsupported role",
		},
		{
			name: "command config missing",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].Executor = ExecutorCommand
				p.Stages[0].Builtin = nil
				p.Stages[0].Command = nil
				return p
			},
			contains: "must declare command",
		},
		{
			name: "artifact missing path",
			build: func() *Pipeline {
				p := validPipelineWithArtifact()
				p.Stages[0].Artifacts[0].Path = ""
				return p
			},
			contains: "must declare path",
		},
		{
			name: "unsupported primary output mode",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].PrimaryOutput = &PrimaryOutputConfig{From: PrimaryOutputFrom("file")}
				return p
			},
			contains: "unsupported primary_output.from",
		},
		{
			name: "final output requires primary output",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].PrimaryOutput = nil
				return p
			},
			contains: "is final_output but has no primary_output",
		},
		{
			name: "must declare exactly one final output stage",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].FinalOutput = false
				return p
			},
			contains: "must declare exactly one final_output stage",
		},
		{
			name: "artifact input unknown stage",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages = append(p.Stages, Stage{
					ID:       "publish",
					Executor: ExecutorBuiltin,
					Builtin:  &BuiltinConfig{Name: "noop"},
					Input: &StageInput{
						From:     StageInputArtifact,
						Stage:    "missing",
						Artifact: "note",
					},
				})
				return p
			},
			contains: `references unknown stage "missing"`,
		},
		{
			name: "artifact input unknown artifact",
			build: func() *Pipeline {
				p := validMinimalPipeline()
				p.Stages[0].Artifacts = []ArtifactDeclaration{{Name: "note", Path: "note.md"}}
				p.Stages = append(p.Stages, Stage{
					ID:       "publish",
					Executor: ExecutorBuiltin,
					Builtin:  &BuiltinConfig{Name: "noop"},
					Input: &StageInput{
						From:     StageInputArtifact,
						Stage:    "render",
						Artifact: "summary",
					},
				})
				return p
			},
			contains: `references unknown artifact "summary" on stage "render"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Validate(tt.build())
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.contains)
		})
	}
}

func TestResolvePipelinePathAndBuiltinSupport(t *testing.T) {
	t.Parallel()

	p := &Pipeline{FilePath: "/tmp/pipelines/demo.yaml"}

	require.Equal(t, "already-absolute", resolvePipelinePath(&Pipeline{}, "already-absolute"))
	require.Equal(t, "/tmp/pipelines/relative/pattern.md", resolvePipelinePath(p, "./relative/pattern.md"))
	require.Equal(t, "/tmp/pipelines/named-pattern", resolvePipelinePath(p, "named-pattern"))
	require.Equal(t, "named-pattern", resolvePipelinePath(&Pipeline{}, "named-pattern"))

	require.True(t, isSupportedBuiltin("passthrough"))
	require.True(t, isSupportedBuiltin("write_publish_manifest"))
	require.False(t, isSupportedBuiltin("imaginary_builtin"))
}

func validMinimalPipeline() *Pipeline {
	return &Pipeline{
		Version:  1,
		Name:     "valid",
		FileStem: "valid",
		Stages: []Stage{
			{
				ID:            "render",
				Executor:      ExecutorBuiltin,
				Builtin:       &BuiltinConfig{Name: "passthrough"},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}
}

func validPipelineWithArtifact() *Pipeline {
	p := validMinimalPipeline()
	p.Stages[0].Artifacts = []ArtifactDeclaration{
		{Name: "note", Path: "note.md"},
	}
	p.Stages[0].PrimaryOutput = &PrimaryOutputConfig{From: PrimaryOutputArtifact, Artifact: "note"}
	return p
}
