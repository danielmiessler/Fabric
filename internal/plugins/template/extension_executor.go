package template

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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
	// Get and verify extension from registry
	ext, err := e.registry.GetExtension(name)
	if err != nil {
		return "", fmt.Errorf("failed to get extension: %w", err)
	}

	// Format the command using our template system
	cmdStr, err := e.formatCommand(ext, operation, value)
	if err != nil {
		return "", fmt.Errorf("failed to format command: %w", err)
	}

	// Split the command string into command and arguments
	cmdParts := strings.Fields(cmdStr)
	if len(cmdParts) < 1 {
		return "", fmt.Errorf("empty command after formatting")
	}

	// Create command with the Executable and formatted arguments
	cmd := exec.Command("sh", "-c", cmdStr)
	//cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	// Set up environment if specified
	if len(ext.Env) > 0 {
		cmd.Env = append(os.Environ(), ext.Env...)
	}

	// Execute based on output method
	outputMethod := ext.GetOutputMethod()
	if outputMethod == "file" {
		return e.executeWithFile(cmd, ext)
	}
	return e.executeStdout(cmd, ext)
}

// formatCommand uses fabric's template system to format the command
// It creates a variables map for the template system using the input values
func (e *ExtensionExecutor) formatCommand(ext *ExtensionDefinition, operation string, value string) (string, error) {
	// Get operation config
	opConfig, exists := ext.Operations[operation]
	if !exists {
		return "", fmt.Errorf("operation %s not found for extension %s", operation, ext.Name)
	}

	vars := make(map[string]string)
	vars["executable"] = ext.Executable
	vars["operation"] = operation
	vars["value"] = value

	// Split on pipe for numbered variables
	values := strings.Split(value, "|")
	for i, val := range values {
		vars[fmt.Sprintf("%d", i+1)] = val
	}

	return ApplyTemplate(opConfig.CmdTemplate, vars, "")
}

// executeStdout runs the command and captures its stdout
func (e *ExtensionExecutor) executeStdout(cmd *exec.Cmd, ext *ExtensionDefinition) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	//debug output
	fmt.Printf("Executing command: %s\n", cmd.String())

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("execution failed: %w\nstderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

// executeWithFile runs the command and handles file-based output
func (e *ExtensionExecutor) executeWithFile(cmd *exec.Cmd, ext *ExtensionDefinition) (string, error) {
	// Parse timeout - this is now a first-class field
	timeout, err := time.ParseDuration(ext.Timeout)
	if err != nil {
		return "", fmt.Errorf("invalid timeout format: %w", err)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// Store the original environment
	originalEnv := cmd.Env
	// Create a new command with context. This might reset Env, depending on the Go version.
	cmd = exec.CommandContext(ctx, cmd.Path, cmd.Args[1:]...)
	// Restore the environment variables explicitly
	cmd.Env = originalEnv

	fileConfig := ext.GetFileConfig()
	if fileConfig == nil {
		return "", fmt.Errorf("no file configuration found")
	}

	// Handle path from stdout case
	if pathFromStdout, ok := fileConfig["path_from_stdout"].(bool); ok && pathFromStdout {
		return e.handlePathFromStdout(cmd, ext)
	}

	// Handle fixed file case
	workDir, _ := fileConfig["work_dir"].(string)
	outputFile, _ := fileConfig["output_file"].(string)

	if outputFile == "" {
		return "", fmt.Errorf("no output file specified in configuration")
	}

	// Set working directory if specified
	if workDir != "" {
		cmd.Dir = workDir
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after %v", timeout)
		}
		return "", fmt.Errorf("execution failed: %w\nerr: %s", err, stderr.String())
	}

	// Construct full file path
	outputPath := outputFile
	if workDir != "" {
		outputPath = filepath.Join(workDir, outputFile)
	}

	content, err := os.ReadFile(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read output file: %w", err)
	}

	// Handle cleanup if enabled
	if ext.IsCleanupEnabled() {
		defer os.Remove(outputPath)
	}

	return string(content), nil
}

// Helper method to handle path from stdout case
func (e *ExtensionExecutor) handlePathFromStdout(cmd *exec.Cmd, ext *ExtensionDefinition) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get output path: %w\nerr: %s", err, stderr.String())
	}

	outputPath := strings.TrimSpace(stdout.String())
	content, err := os.ReadFile(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read output file: %w", err)
	}

	if ext.IsCleanupEnabled() {
		defer os.Remove(outputPath)
	}

	return string(content), nil
}
