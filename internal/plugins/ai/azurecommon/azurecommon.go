package azurecommon

import (
	"strings"
)

// DefaultAPIVersion is the default Azure OpenAI API version.
const DefaultAPIVersion = "2025-04-01-preview"

// ParseDeployments splits a comma-separated deployment string into a slice,
// trimming whitespace and discarding empty entries.
func ParseDeployments(value string) []string {
	parts := strings.Split(value, ",")
	var deployments []string
	for _, part := range parts {
		if deployment := strings.TrimSpace(part); deployment != "" {
			deployments = append(deployments, deployment)
		}
	}
	return deployments
}
