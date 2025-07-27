package changelog

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const DefaultSummarizeModel = "claude-sonnet-4-20250514"
const MinContentLength = 256 // Minimum content length to consider for summarization

const prompt = `# ROLE
You are an expert Technical Writer specializing in creating clear, concise,
and professional release notes from raw Git commit logs.

# TASK
Your goal is to transform a provided block of Git commit logs into a clean,
human-readable changelog summary. You will identify the most important changes,
format them as a bulleted list, and preserve the associated Pull Request (PR)
information.

# INSTRUCTIONS:
Follow these steps in order:
1. Deeply analyze the input. You will be given a block of text containing PR
   information and commit log messages. Carefully read through the logs
   to identify individual commits and their descriptions.
2. Identify Key Changes: Focus on commits that represent significant changes,
   such as new features ("feat"), bug fixes ("fix"), performance improvements ("perf"),
   or breaking changes ("BREAKING CHANGE").
3. Select the Top 5: From the identified key changes, select a maximum of the five
   (5) most impactful ones to include in the summary.
   If there are five or fewer total changes, include all of them.
4. Format the Output:
    - Where you see a PR header, include the PR header verbatim. NO CHANGES.
	  **This is a critical rule: Do not modify the PR header, as it contains
	  important links.** What follow the PR header are the related changes.
	- Do not add any additional text or preamble. Begin directly with the output.
	- Use bullet points for each key change. Starting each point with a hyphen ("-").
	- Ensure that the summary is concise and focused on the main changes.
	- The summary should be in American English (en-US), using proper grammar and punctuation.
5. If the content is too brief or you do not see any PR headers, return the content as is.
`

// getSummarizeModel returns the model to use for AI summarization
func getSummarizeModel() string {
	if model := os.Getenv("FABRIC_CHANGELOG_SUMMARIZE_MODEL"); model != "" {
		return model
	}
	return DefaultSummarizeModel
}

// SummarizeVersionContent takes raw version content and returns AI-enhanced summary
func SummarizeVersionContent(content string) (string, error) {
	if strings.TrimSpace(content) == "" {
		return "", fmt.Errorf("no content to summarize")
	}
	if len(content) < MinContentLength {
		// If content is too brief, return it as is
		return content, nil
	}

	model := getSummarizeModel()

	cmd := exec.Command("fabric", "-m", model, prompt)
	cmd.Stdin = strings.NewReader(content)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("fabric command failed: %w", err)
	}

	summary := strings.TrimSpace(string(output))
	if summary == "" {
		return "", fmt.Errorf("fabric returned empty summary")
	}

	return summary, nil
}
