package domain

// WorkflowDefinition represents a YAML workflow file with sequential pattern steps.
type WorkflowDefinition struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Steps       []WorkflowStep `yaml:"steps"`
}

// WorkflowStep represents a single step in a workflow pipeline.
type WorkflowStep struct {
	Pattern   string            `yaml:"pattern"`
	Variables map[string]string `yaml:"variables"`
}
