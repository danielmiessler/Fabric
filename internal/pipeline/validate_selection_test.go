package pipeline

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateRejectsInvalidPipelineContracts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mutate  func(*Pipeline)
		wantErr string
	}{
		{
			name: "unsupported accepted source mode",
			mutate: func(p *Pipeline) {
				p.Accepts = []SourceMode{SourceMode("directory")}
			},
			wantErr: "unsupported source mode",
		},
		{
			name: "builtin stage requires name",
			mutate: func(p *Pipeline) {
				p.Stages[0].Builtin = &BuiltinConfig{}
			},
			wantErr: "must declare builtin.name",
		},
		{
			name: "command stage requires program",
			mutate: func(p *Pipeline) {
				p.Stages[0].Executor = ExecutorCommand
				p.Stages[0].Builtin = nil
				p.Stages[0].Command = &CommandConfig{}
			},
			wantErr: "must declare command.program",
		},
		{
			name: "command stage rejects negative timeout",
			mutate: func(p *Pipeline) {
				p.Stages[0].Executor = ExecutorCommand
				p.Stages[0].Builtin = nil
				p.Stages[0].Command = &CommandConfig{Program: "echo", Timeout: -1}
			},
			wantErr: "timeout must be >= 0",
		},
		{
			name: "pattern stage requires pattern",
			mutate: func(p *Pipeline) {
				p.Stages[0].Executor = ExecutorFabricPattern
				p.Stages[0].Builtin = nil
				p.Stages[0].Pattern = ""
			},
			wantErr: "must declare pattern",
		},
		{
			name: "rejects duplicate artifact names",
			mutate: func(p *Pipeline) {
				p.Stages[0].Artifacts = []ArtifactDeclaration{
					{Name: "note", Path: "note.md"},
					{Name: "note", Path: "copy.md"},
				}
			},
			wantErr: "duplicate artifact name",
		},
		{
			name: "rejects absolute artifact path",
			mutate: func(p *Pipeline) {
				p.Stages[0].Artifacts = []ArtifactDeclaration{
					{Name: "note", Path: "/tmp/note.md"},
				}
			},
			wantErr: "path must be relative",
		},
		{
			name: "artifact input requires stage",
			mutate: func(p *Pipeline) {
				p.Stages = append(p.Stages, Stage{
					ID:       "reader",
					Executor: ExecutorBuiltin,
					Builtin:  &BuiltinConfig{Name: "passthrough"},
					Input: &StageInput{
						From:     StageInputArtifact,
						Artifact: "note",
					},
				})
			},
			wantErr: "requires input.stage",
		},
		{
			name: "artifact input requires artifact name",
			mutate: func(p *Pipeline) {
				p.Stages = append(p.Stages, Stage{
					ID:       "reader",
					Executor: ExecutorBuiltin,
					Builtin:  &BuiltinConfig{Name: "passthrough"},
					Input: &StageInput{
						From:  StageInputArtifact,
						Stage: "render",
					},
				})
			},
			wantErr: "requires input.artifact",
		},
		{
			name: "unsupported input source rejected",
			mutate: func(p *Pipeline) {
				p.Stages[0].Input = &StageInput{From: StageInputFrom("future")}
			},
			wantErr: "unsupported input.from",
		},
		{
			name: "primary output artifact must be declared",
			mutate: func(p *Pipeline) {
				p.Stages[0].PrimaryOutput = &PrimaryOutputConfig{
					From:     PrimaryOutputArtifact,
					Artifact: "missing",
				}
			},
			wantErr: "primary_output artifact",
		},
		{
			name: "final output requires primary output",
			mutate: func(p *Pipeline) {
				p.Stages[0].PrimaryOutput = nil
			},
			wantErr: "is final_output but has no primary_output",
		},
		{
			name: "artifact input requires known source stage",
			mutate: func(p *Pipeline) {
				p.Stages = append(p.Stages, Stage{
					ID:       "reader",
					Executor: ExecutorBuiltin,
					Builtin:  &BuiltinConfig{Name: "passthrough"},
					Input: &StageInput{
						From:     StageInputArtifact,
						Stage:    "missing",
						Artifact: "note",
					},
				})
			},
			wantErr: "references unknown stage",
		},
		{
			name: "artifact input requires declared artifact",
			mutate: func(p *Pipeline) {
				p.Stages = append(p.Stages, Stage{
					ID:       "reader",
					Executor: ExecutorBuiltin,
					Builtin:  &BuiltinConfig{Name: "passthrough"},
					Input: &StageInput{
						From:     StageInputArtifact,
						Stage:    "render",
						Artifact: "missing",
					},
				})
			},
			wantErr: "references unknown artifact",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := validPipelineForContractTests()
			tt.mutate(p)

			err := Validate(p)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestValidateAcceptedSource(t *testing.T) {
	t.Parallel()

	p := validPipelineForContractTests()
	p.Accepts = []SourceMode{SourceModeStdin, SourceModeScrapeURL}

	require.NoError(t, validateAcceptedSource(p, SourceModeStdin))
	require.NoError(t, validateAcceptedSource(&Pipeline{Name: "implicit-any"}, SourceModeSource))

	err := validateAcceptedSource(p, SourceModeSource)
	require.Error(t, err)
	require.Contains(t, err.Error(), `does not accept source mode "source"`)
}

func TestSelectStageIndices(t *testing.T) {
	t.Parallel()

	stages := []Stage{
		{ID: "ingest"},
		{ID: "render"},
		{ID: "validate"},
		{ID: "publish"},
	}

	indices, set, err := SelectStageIndices(stages, "render", "validate", "")
	require.NoError(t, err)
	require.Equal(t, []int{1, 2}, indices)
	require.Contains(t, set, 1)
	require.Contains(t, set, 2)

	indices, set, err = SelectStageIndices(stages, "", "", "publish")
	require.NoError(t, err)
	require.Equal(t, []int{3}, indices)
	require.Equal(t, map[int]struct{}{3: {}}, set)

	_, _, err = SelectStageIndices(nil, "", "", "")
	require.Error(t, err)
	require.Contains(t, err.Error(), "pipeline has no stages")

	_, _, err = SelectStageIndices(stages, "render", "", "publish")
	require.Error(t, err)
	require.Contains(t, err.Error(), "--only-stage cannot be combined")

	_, _, err = SelectStageIndices(stages, "missing", "", "")
	require.Error(t, err)
	require.Contains(t, err.Error(), `stage "missing" not found`)

	_, _, err = SelectStageIndices(stages, "", "missing", "")
	require.Error(t, err)
	require.Contains(t, err.Error(), `stage "missing" not found`)

	_, _, err = SelectStageIndices(stages, "", "", "missing")
	require.Error(t, err)
	require.Contains(t, err.Error(), `stage "missing" not found`)

	_, _, err = SelectStageIndices(stages, "validate", "render", "")
	require.Error(t, err)
	require.Contains(t, err.Error(), "--from-stage")
}

func validPipelineForContractTests() *Pipeline {
	return &Pipeline{
		Version:  1,
		Name:     "contract-checks",
		FileStem: "contract-checks",
		Stages: []Stage{
			{
				ID:       "render",
				Executor: ExecutorBuiltin,
				Builtin:  &BuiltinConfig{Name: "passthrough"},
				Artifacts: []ArtifactDeclaration{
					{Name: "note", Path: "note.md", Required: boolPtr(false)},
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}
}
