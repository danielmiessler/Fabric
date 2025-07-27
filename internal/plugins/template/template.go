package template

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	textPlugin     = &TextPlugin{}
	datetimePlugin = &DateTimePlugin{}
	filePlugin     = &FilePlugin{}
	fetchPlugin    = &FetchPlugin{}
	sysPlugin      = &SysPlugin{}
	Debug          = false // Debug flag
)

var extensionManager *ExtensionManager

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		debugf("Warning: could not initialize extension manager: %v\n", err)
	}
	configDir := filepath.Join(homedir, ".config/fabric")
	extensionManager = NewExtensionManager(configDir)
	// Extensions will work if registry exists, otherwise they'll just fail gracefully
}

var pluginPattern = regexp.MustCompile(`\{\{plugin:([^:]+):([^:]+)(?::([^}]+))?\}\}`)
var extensionPattern = regexp.MustCompile(`\{\{ext:([^:]+):([^:]+)(?::([^}]+))?\}\}`)

func debugf(format string, a ...interface{}) {
	if Debug {
		fmt.Printf(format, a...)
	}
}

func ApplyTemplate(content string, variables map[string]string, input string) (string, error) {

	var missingVars []string
	r := regexp.MustCompile(`\{\{([^{}]+)\}\}`)

	debugf("Starting template processing\n")
	for strings.Contains(content, "{{") {
		matches := r.FindAllStringSubmatch(content, -1)
		if len(matches) == 0 {
			break
		}

		replaced := false
		for _, match := range matches {
			fullMatch := match[0]
			varName := match[1]

			// Check if this is a plugin call
			if strings.HasPrefix(varName, "plugin:") {
				pluginMatches := pluginPattern.FindStringSubmatch(fullMatch)
				if len(pluginMatches) >= 3 {
					namespace := pluginMatches[1]
					operation := pluginMatches[2]
					value := ""
					if len(pluginMatches) == 4 {
						value = pluginMatches[3]
					}

					debugf("\nPlugin call:\n")
					debugf("  Namespace: %s\n", namespace)
					debugf("  Operation: %s\n", operation)
					debugf("  Value: %s\n", value)

					var result string
					var err error

					switch namespace {
					case "text":
						debugf("Executing text plugin\n")
						result, err = textPlugin.Apply(operation, value)
					case "datetime":
						debugf("Executing datetime plugin\n")
						result, err = datetimePlugin.Apply(operation, value)
					case "file":
						debugf("Executing file plugin\n")
						result, err = filePlugin.Apply(operation, value)
						debugf("File plugin result: %#v\n", result)
					case "fetch":
						debugf("Executing fetch plugin\n")
						result, err = fetchPlugin.Apply(operation, value)
					case "sys":
						debugf("Executing sys plugin\n")
						result, err = sysPlugin.Apply(operation, value)
					default:
						return "", fmt.Errorf("unknown plugin namespace: %s", namespace)
					}

					if err != nil {
						debugf("Plugin error: %v\n", err)
						return "", fmt.Errorf("plugin %s error: %v", namespace, err)
					}

					debugf("Plugin result: %s\n", result)
					content = strings.ReplaceAll(content, fullMatch, result)
					debugf("Content after replacement: %s\n", content)
					continue
				}
			}

			if pluginMatches := extensionPattern.FindStringSubmatch(fullMatch); len(pluginMatches) >= 3 {
				name := pluginMatches[1]
				operation := pluginMatches[2]
				value := ""
				if len(pluginMatches) == 4 {
					value = pluginMatches[3]
				}

				debugf("\nExtension call:\n")
				debugf("  Name: %s\n", name)
				debugf("  Operation: %s\n", operation)
				debugf("  Value: %s\n", value)

				result, err := extensionManager.ProcessExtension(name, operation, value)
				if err != nil {
					return "", fmt.Errorf("extension %s error: %v", name, err)
				}

				content = strings.ReplaceAll(content, fullMatch, result)
				replaced = true
				continue
			}

			// Handle regular variables and input
			debugf("Processing variable: %s\n", varName)
			if varName == "input" {
				debugf("Replacing {{input}}\n")
				replaced = true
				content = strings.ReplaceAll(content, fullMatch, input)
			} else {
				if val, ok := variables[varName]; !ok {
					debugf("Missing variable: %s\n", varName)
					missingVars = append(missingVars, varName)
					return "", fmt.Errorf("missing required variable: %s", varName)
				} else {
					debugf("Replacing variable %s with value: %s\n", varName, val)
					content = strings.ReplaceAll(content, fullMatch, val)
					replaced = true
				}
			}
			if !replaced {
				return "", fmt.Errorf("template processing stuck - potential infinite loop")
			}
		}
	}

	debugf("Starting template processing\n")
	for strings.Contains(content, "{{") {
		matches := r.FindAllStringSubmatch(content, -1)
		if len(matches) == 0 {
			break
		}

		replaced := false
		for _, match := range matches {
			fullMatch := match[0]
			varName := match[1]

			// Check if this is a plugin call
			if strings.HasPrefix(varName, "plugin:") {
				pluginMatches := pluginPattern.FindStringSubmatch(fullMatch)
				if len(pluginMatches) >= 3 {
					namespace := pluginMatches[1]
					operation := pluginMatches[2]
					value := ""
					if len(pluginMatches) == 4 {
						value = pluginMatches[3]
					}

					debugf("\nPlugin call:\n")
					debugf("  Namespace: %s\n", namespace)
					debugf("  Operation: %s\n", operation)
					debugf("  Value: %s\n", value)

					var result string
					var err error

					switch namespace {
					case "text":
						debugf("Executing text plugin\n")
						result, err = textPlugin.Apply(operation, value)
					case "datetime":
						debugf("Executing datetime plugin\n")
						result, err = datetimePlugin.Apply(operation, value)
					case "file":
						debugf("Executing file plugin\n")
						result, err = filePlugin.Apply(operation, value)
						debugf("File plugin result: %#v\n", result)
					case "fetch":
						debugf("Executing fetch plugin\n")
						result, err = fetchPlugin.Apply(operation, value)
					case "sys":
						debugf("Executing sys plugin\n")
						result, err = sysPlugin.Apply(operation, value)
					default:
						return "", fmt.Errorf("unknown plugin namespace: %s", namespace)
					}

					if err != nil {
						debugf("Plugin error: %v\n", err)
						return "", fmt.Errorf("plugin %s error: %v", namespace, err)
					}

					debugf("Plugin result: %s\n", result)
					content = strings.ReplaceAll(content, fullMatch, result)
					debugf("Content after replacement: %s\n", content)
					continue
				}
			}

			// Handle regular variables and input
			debugf("Processing variable: %s\n", varName)
			if varName == "input" {
				debugf("Replacing {{input}}\n")
				replaced = true
				content = strings.ReplaceAll(content, fullMatch, input)
			} else {
				if val, ok := variables[varName]; !ok {
					debugf("Missing variable: %s\n", varName)
					missingVars = append(missingVars, varName)
					return "", fmt.Errorf("missing required variable: %s", varName)
				} else {
					debugf("Replacing variable %s with value: %s\n", varName, val)
					content = strings.ReplaceAll(content, fullMatch, val)
					replaced = true
				}
			}
			if !replaced {
				return "", fmt.Errorf("template processing stuck - potential infinite loop")
			}
		}
	}

	debugf("Template processing complete\n")
	return content, nil
}
