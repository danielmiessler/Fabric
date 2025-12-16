package template

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	debuglog "github.com/danielmiessler/fabric/internal/log"
)

var (
	textPlugin     = &TextPlugin{}
	datetimePlugin = &DateTimePlugin{}
	filePlugin     = &FilePlugin{}
	fetchPlugin    = &FetchPlugin{}
	sysPlugin      = &SysPlugin{}
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

func debugf(format string, a ...any) {
	debuglog.Debug(debuglog.Trace, format, a...)
}

// matchTriple extracts the first two required and optional third value from a token
// pattern of the form {{type:part1:part2(:part3)?}} returning part1, part2, part3 (possibly empty)
func matchTriple(r *regexp.Regexp, full string) (string, string, string, bool) {
	parts := r.FindStringSubmatch(full)
	if len(parts) >= 3 {
		v := ""
		if len(parts) == 4 {
			v = parts[3]
		}
		return parts[1], parts[2], v, true
	}
	return "", "", "", false
}

func ApplyTemplate(content string, variables map[string]string, input string) (string, error) {
	tokenPattern := regexp.MustCompile(`\{\{([^{}]+)\}\}`)

	debugf("Starting template processing with input='%s'\n", input)

	for {
		if !strings.Contains(content, "{{") {
			break
		}
		matches := tokenPattern.FindAllStringSubmatch(content, -1)
		if len(matches) == 0 {
			break
		}

		progress := false
		for _, m := range matches {
			full := m[0]
			raw := m[1]

			// Extension call
			if strings.HasPrefix(raw, "ext:") {
				if name, operation, value, ok := matchTriple(extensionPattern, full); ok {
					if strings.Contains(value, InputSentinel) {
						value = strings.ReplaceAll(value, InputSentinel, input)
						debugf("Replaced sentinel in extension value with input\n")
					}
					debugf("Extension call: name=%s operation=%s value=%s\n", name, operation, value)
					result, err := extensionManager.ProcessExtension(name, operation, value)
					if err != nil {
						return "", fmt.Errorf("extension %s error: %v", name, err)
					}
					content = strings.ReplaceAll(content, full, result)
					progress = true
					continue
				}
			}

			// Plugin call
			if strings.HasPrefix(raw, "plugin:") {
				if namespace, operation, value, ok := matchTriple(pluginPattern, full); ok {
					debugf("Plugin call: namespace=%s operation=%s value=%s\n", namespace, operation, value)
					var (
						result string
						err    error
					)
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
					content = strings.ReplaceAll(content, full, result)
					progress = true
					continue
				}
			}

			// Variables / input / sentinel
			switch raw {
			case "input", InputSentinel:
				content = strings.ReplaceAll(content, full, input)
				progress = true
			default:
				val, ok := variables[raw]
				if !ok {
					return "", fmt.Errorf("missing required variable: %s", raw)
				}
				content = strings.ReplaceAll(content, full, val)
				progress = true
			}
		}

		if !progress {
			return "", fmt.Errorf("template processing stuck - potential infinite loop")
		}
	}

	debugf("Template processing complete\n")
	return content, nil
}
