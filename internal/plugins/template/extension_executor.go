package template

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// ExtensionExecutor handles the secure execution of extensions
// It uses the registry to verify extensions before running them
type ExtensionExecutor struct {
	registry *ExtensionRegistry
}

// NewExtensionExecutor creates a new executor instance
// It requires a registry to verify extensions
func NewExtensionExecutor(registry *ExtensionRegistry) *ExtensionExecutor {
	return &ExtensionExecutor{
		registry: registry,
	}
}

// Execute runs an extension with the given operation and value string
// name: the registered name of the extension
// operation: the operation to perform
// value: the input value(s) for the operation
// In extension_executor.go
func (e *ExtensionExecutor) Execute(name, operation, value string) (string, error) {
	ext, err := e.registry.GetExtension(name)
	if err != nil {
		return "", fmt.Errorf(i18n.T("extension_failed_get_extension"), err)
	}

	cmdStr, err := e.formatCommand(ext, operation, value)
	if err != nil {
		return "", fmt.Errorf(i18n.T("extension_failed_format_command"), err)
	}

	cmdParts := strings.Fields(cmdStr)
	if len(cmdParts) < 1 {
		return "", errors.New(i18n.T("extension_empty_command"))
	}

	cmd := exec.Command("sh", "-c", cmdStr)

	if len(ext.Env) > 0 {
		cmd.Env = append(os.Environ(), ext.Env...)
	}

	outputMethod := ext.GetOutputMethod()
	if outputMethod == "file" {
		return e.executeWithFile(cmd, ext)
	}
	return e.executeStdout(cmd, ext)
}

// formatCommand fills the operation's cmd_template with ApplyTemplate.
func (e *ExtensionExecutor) formatCommand(ext *ExtensionDefinition, operation string, value string) (string, error) {
	opConfig, exists := ext.Operations[operation]
	if !exists {
		return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_operation_not_found"), operation, ext.Name))
	}

	// Shell-escape every user-controlled value. Execute passes the command
	// string to "sh -c", so shell metacharacters in a raw value would run as
	// commands. Single quotes make each value one literal argument.
	vars := make(map[string]string)
	vars["executable"] = ext.Executable
	vars["operation"] = operation
	vars["value"] = shellEscape(value)

	// Split on pipe for numbered variables
	values := strings.Split(value, "|")
	for i, val := range values {
		vars[fmt.Sprintf("%d", i+1)] = shellEscape(val)
	}

	return ApplyTemplate(opConfig.CmdTemplate, vars, "")
}

// shellEscape wraps s in single quotes and escapes embedded single quotes.
// The result is one literal argument to "sh -c".
func shellEscape(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
}

// executeStdout runs the command and captures its stdout
func (e *ExtensionExecutor) executeStdout(cmd *exec.Cmd, ext *ExtensionDefinition) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Printf(i18n.T("extension_executing_command"), cmd.String())

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(i18n.T("extension_execution_failed_stderr"), err, stderr.String())
	}

	return stdout.String(), nil
}

// executeWithFile runs the command and reads the result from its output file.
func (e *ExtensionExecutor) executeWithFile(cmd *exec.Cmd, ext *ExtensionDefinition) (string, error) {
	timeout, err := time.ParseDuration(ext.Timeout)
	if err != nil {
		return "", fmt.Errorf(i18n.T("extension_invalid_timeout_format"), err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// exec.CommandContext returns a new Cmd, so copy Env across.
	originalEnv := cmd.Env
	cmd = exec.CommandContext(ctx, cmd.Path, cmd.Args[1:]...)
	cmd.Env = originalEnv

	fileConfig := ext.GetFileConfig()
	if fileConfig == nil {
		return "", errors.New(i18n.T("extension_no_file_config"))
	}

	if pathFromStdout, ok := fileConfig["path_from_stdout"].(bool); ok && pathFromStdout {
		return e.handlePathFromStdout(cmd, ext)
	}

	workDir, _ := fileConfig["work_dir"].(string)
	outputFile, _ := fileConfig["output_file"].(string)

	if outputFile == "" {
		return "", errors.New(i18n.T("extension_no_output_file"))
	}

	if workDir != "" {
		cmd.Dir = workDir
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_execution_timed_out"), timeout))
		}
		return "", fmt.Errorf(i18n.T("extension_execution_failed_err"), err, stderr.String())
	}

	outputPath := outputFile
	if workDir != "" {
		outputPath = filepath.Join(workDir, outputFile)
	}

	content, err := os.ReadFile(outputPath)
	if err != nil {
		return "", fmt.Errorf(i18n.T("extension_failed_read_output_file"), err)
	}

	if ext.IsCleanupEnabled() {
		defer os.Remove(outputPath)
	}

	return string(content), nil
}

// handlePathFromStdout runs the command and reads the file whose path the command prints to stdout.
func (e *ExtensionExecutor) handlePathFromStdout(cmd *exec.Cmd, ext *ExtensionDefinition) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(i18n.T("extension_failed_get_output_path"), err, stderr.String())
	}

	outputPath := strings.TrimSpace(stdout.String())
	content, err := os.ReadFile(outputPath)
	if err != nil {
		return "", fmt.Errorf(i18n.T("extension_failed_read_output_file"), err)
	}

	if ext.IsCleanupEnabled() {
		defer os.Remove(outputPath)
	}

	return string(content), nil
}
