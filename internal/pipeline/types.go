package pipeline

import "time"

type SourceMode string

const (
	SourceModeStdin     SourceMode = "stdin"
	SourceModeSource    SourceMode = "source"
	SourceModeScrapeURL SourceMode = "scrape_url"
)

type ExecutorType string

const (
	ExecutorCommand       ExecutorType = "command"
	ExecutorFabricPattern ExecutorType = "fabric_pattern"
	ExecutorBuiltin       ExecutorType = "builtin"
)

type StageRole string

const (
	StageRoleDefault  StageRole = ""
	StageRoleValidate StageRole = "validate"
	StageRolePublish  StageRole = "publish"
)

type StageInputFrom string

const (
	StageInputSource   StageInputFrom = "source"
	StageInputPrevious StageInputFrom = "previous"
	StageInputArtifact StageInputFrom = "artifact"
)

type PrimaryOutputFrom string

const (
	PrimaryOutputStdout   PrimaryOutputFrom = "stdout"
	PrimaryOutputArtifact PrimaryOutputFrom = "artifact"
)

type DefinitionSource string

const (
	DefinitionSourceBuiltIn DefinitionSource = "built-in"
	DefinitionSourceUser    DefinitionSource = "user"
)

type Pipeline struct {
	Version     int          `yaml:"version"`
	Name        string       `yaml:"name"`
	Description string       `yaml:"description,omitempty"`
	Accepts     []SourceMode `yaml:"accepts,omitempty"`
	Stages      []Stage      `yaml:"stages"`

	FilePath         string           `yaml:"-"`
	FileName         string           `yaml:"-"`
	FileStem         string           `yaml:"-"`
	DefinitionSource DefinitionSource `yaml:"-"`
}

type Stage struct {
	ID            string                `yaml:"id"`
	Role          StageRole             `yaml:"role,omitempty"`
	Executor      ExecutorType          `yaml:"executor"`
	Input         *StageInput           `yaml:"input,omitempty"`
	Command       *CommandConfig        `yaml:"command,omitempty"`
	Pattern       string                `yaml:"pattern,omitempty"`
	Context       string                `yaml:"context,omitempty"`
	Strategy      string                `yaml:"strategy,omitempty"`
	Variables     map[string]string     `yaml:"variables,omitempty"`
	Stream        bool                  `yaml:"stream,omitempty"`
	Builtin       *BuiltinConfig        `yaml:"builtin,omitempty"`
	Artifacts     []ArtifactDeclaration `yaml:"artifacts,omitempty"`
	PrimaryOutput *PrimaryOutputConfig  `yaml:"primary_output,omitempty"`
	FinalOutput   bool                  `yaml:"final_output,omitempty"`
}

type StageInput struct {
	From     StageInputFrom `yaml:"from"`
	Stage    string         `yaml:"stage,omitempty"`
	Artifact string         `yaml:"artifact,omitempty"`
}

type CommandConfig struct {
	Program string            `yaml:"program"`
	Args    []string          `yaml:"args,omitempty"`
	Cwd     string            `yaml:"cwd,omitempty"`
	Env     map[string]string `yaml:"env,omitempty"`
	Timeout int               `yaml:"timeout,omitempty"`
}

type BuiltinConfig struct {
	Name   string                 `yaml:"name"`
	Config map[string]interface{} `yaml:"config,omitempty"`
}

type ArtifactDeclaration struct {
	Name     string `yaml:"name"`
	Path     string `yaml:"path"`
	Required *bool  `yaml:"required,omitempty"`
}

func (a ArtifactDeclaration) IsRequired() bool {
	return a.Required == nil || *a.Required
}

type PrimaryOutputConfig struct {
	From     PrimaryOutputFrom `yaml:"from"`
	Artifact string            `yaml:"artifact,omitempty"`
}

type DiscoveryEntry struct {
	Name             string
	Path             string
	DefinitionSource DefinitionSource
	OverridesBuiltIn bool
}

type RunSource struct {
	Mode      SourceMode `json:"mode"`
	Reference string     `json:"reference,omitempty"`
	Payload   string     `json:"-"`
}

type RunManifest struct {
	RunID        string             `json:"run_id"`
	PipelineName string             `json:"pipeline_name"`
	PipelineFile string             `json:"pipeline_file"`
	Status       string             `json:"status"`
	StartedAt    time.Time          `json:"started_at"`
	FinishedAt   *time.Time         `json:"finished_at,omitempty"`
	Source       RunSourceManifest  `json:"source"`
	Stages       []RunStageManifest `json:"stages"`
	Warnings     []string           `json:"warnings,omitempty"`
	FinalOutput  *FinalOutputReport `json:"final_output,omitempty"`
}

type RunSourceManifest struct {
	Mode      SourceMode `json:"mode"`
	Reference string     `json:"reference,omitempty"`
}

type RunStageManifest struct {
	ID         string       `json:"id"`
	Role       StageRole    `json:"role,omitempty"`
	Executor   ExecutorType `json:"executor"`
	Status     string       `json:"status"`
	Error      string       `json:"error,omitempty"`
	StartedAt  *time.Time   `json:"started_at,omitempty"`
	FinishedAt *time.Time   `json:"finished_at,omitempty"`
	Files      []string     `json:"files,omitempty"`
}

type SourceManifest struct {
	Mode         SourceMode `json:"mode"`
	Reference    string     `json:"reference,omitempty"`
	PayloadBytes int        `json:"payload_bytes"`
}

type FinalOutputReport struct {
	StageID string `json:"stage_id"`
	Bytes   int    `json:"bytes"`
}
