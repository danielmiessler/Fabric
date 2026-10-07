package template

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
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
	// A missing registry file is not fatal. Extension calls then fail at lookup.
}

var pluginPattern = regexp.MustCompile(`\{\{plugin:([^:]+):([^:]+)(?::([^}]+))?\}\}`)
var extensionPattern = regexp.MustCompile(`\{\{ext:([^:]+):([^:]+)(?::([^}]+))?\}\}`)

func debugf(format string, a ...any) {
	debuglog.Debug(debuglog.Trace, format, a...)
}

// matchTriple matches full against r, a pattern of the form {{type:part1:part2(:part3)?}},
// and returns part1, part2, and part3. part3 is empty when absent.
func matchTriple(r *regexp.Regexp, full string) (string, string, string, bool) {
	// The match must start at the first character of full and stop at the
	// last character. Without this check, a variable value in a nested token
	// can close the outer token and start a new token. An example is the value
	// "}}{{plugin:sys:env:NAME" in "{{plugin:text:upper:{{name}}}}". With this
	// check, such a token goes to the variable lookup, which gives an error.
	if r.FindString(full) != full {
		return "", "", "", false
	}
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

// ApplyTemplate resolves each {{...}} token in content one time, in one pass
// from left to right. It writes the result of a token as literal text and
// does not scan that text again. Thus a variable value that contains
// "{{plugin:sys:env:NAME}}" stays as literal text, and the plugin does not run.
//
// A token in the template can contain other tokens, for example
// "{{plugin:text:upper:{{name}}}}" or "{{ext:name:op:{{input}}}}". The inner
// tokens resolve first, because they are part of the template text.
func ApplyTemplate(content string, variables map[string]string, input string) (string, error) {
	return applyTemplate(content, variables, input, allTokens)
}

// ApplyTemplateNoSystemPlugins is ApplyTemplate for a template from an
// untrusted client, for example a pattern that the REST API saved. It runs
// only the text and datetime plugins. A sys, file or fetch plugin token or
// an extension token gives an error.
func ApplyTemplateNoSystemPlugins(content string, variables map[string]string, input string) (string, error) {
	return applyTemplate(content, variables, input, noSystemPlugins)
}

// ApplyTemplateInput replaces the variables in user input. A plugin or
// extension token in the input stays as literal text, and the plugin or
// extension does not run. The input can contain text from other sources, for
// example a web page or a transcript. The token {{input}} gives an empty
// string.
func ApplyTemplateInput(content string, variables map[string]string) (string, error) {
	return applyTemplate(content, variables, "", varsOnly)
}

// tokenMode tells applyTemplate which tokens it can run.
type tokenMode int

const (
	allTokens       tokenMode = iota // all plugins and extensions
	noSystemPlugins                  // only the text and datetime plugins
	varsOnly                         // no plugin or extension
)

// applyTemplate does the work for ApplyTemplate, ApplyTemplateNoSystemPlugins
// and ApplyTemplateInput.
func applyTemplate(content string, variables map[string]string, input string, mode tokenMode) (string, error) {
	debugf("Starting template processing with input='%s'\n", input)

	// out is the result. open has an entry for each "{{" that has no "}}"
	// yet. A "}}" closes the last open "{{". The loop reads only content,
	// thus it does not scan the text that a token gives. A "{{" that has no
	// "}}" stays in out as literal text.
	type openToken struct {
		out     int  // index of the "{{" in out
		src     int  // index of the "{{" in content
		literal bool // the body has template text that is not an inner token
	}
	var out []byte
	var open []openToken
	markLiteral := func() {
		if len(open) > 0 {
			open[len(open)-1].literal = true
		}
	}
	for i := 0; i < len(content); {
		switch {
		case strings.HasPrefix(content[i:], "{{{") && !strings.HasPrefix(content[i:], "{{{{"):
			// Three braces, as in the Mustache form "{{{name}}}". The first
			// "{" is literal text, and the token starts at the next character.
			markLiteral()
			out = append(out, '{')
			i++
		case strings.HasPrefix(content[i:], "{{"):
			open = append(open, openToken{out: len(out), src: i})
			out = append(out, "{{"...)
			i += 2
		case strings.HasPrefix(content[i:], "}}") && len(open) > 0:
			tok := open[len(open)-1]
			open = open[:len(open)-1]
			src := content[tok.src+2 : i]
			i += 2
			body := string(out[tok.out+2:])
			if body == "" {
				// An empty token ("{{}}") stays as literal text, as before.
				markLiteral()
				out = append(out, "}}"...)
				continue
			}
			if !tok.literal {
				// A body that has only the results of inner tokens, as in
				// "{{{{name}}}}", stays as literal text. Thus a value cannot
				// become a token.
				out = append(out, "}}"...)
				continue
			}
			result, err := resolveToken(body, src, variables, input, mode)
			if err != nil {
				return "", err
			}
			out = append(out[:tok.out], result...)
		default:
			markLiteral()
			out = append(out, content[i])
			i++
		}
	}

	debugf("Template processing complete\n")
	return string(out), nil
}

// resolveToken gives the text for one token body. The body has no outer
// braces, and its nested tokens are already resolved. src is the body as it
// is in the template, before the inner tokens resolve. An error message shows
// src, because body can contain a secret from an inner token. The caller
// writes the result as literal text and does not scan it again.
func resolveToken(body, src string, variables map[string]string, input string, mode tokenMode) (string, error) {
	full := "{{" + body + "}}"

	// With varsOnly, a plugin or extension token stays as literal text. The
	// body contains the text of its inner tokens. Thus an inner token cannot
	// make a plugin or extension token that runs.
	if mode == varsOnly && (strings.HasPrefix(body, "ext:") || strings.HasPrefix(body, "plugin:")) {
		return full, nil
	}

	if strings.HasPrefix(body, "ext:") {
		if name, operation, value, ok := matchTriple(extensionPattern, full); ok {
			if mode == noSystemPlugins {
				return "", fmt.Errorf(i18n.T("template_plugin_not_permitted"), "ext:"+name)
			}
			if strings.Contains(value, InputSentinel) {
				value = strings.ReplaceAll(value, InputSentinel, input)
				debugf("Replaced sentinel in extension value with input\n")
			}
			debugf("Extension call: name=%s operation=%s value=%s\n", name, operation, value)
			result, err := extensionManager.ProcessExtension(name, operation, value)
			if err != nil {
				return "", fmt.Errorf(i18n.T("template_extension_error"), name, err)
			}
			return result, nil
		}
	}

	if strings.HasPrefix(body, "plugin:") {
		if namespace, operation, value, ok := matchTriple(pluginPattern, full); ok {
			if mode == noSystemPlugins && namespace != "text" && namespace != "datetime" {
				return "", fmt.Errorf(i18n.T("template_plugin_not_permitted"), "plugin:"+namespace)
			}
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
				return "", fmt.Errorf(i18n.T("template_unknown_plugin_namespace"), namespace)
			}
			if err != nil {
				debugf("Plugin error: %v\n", err)
				return "", fmt.Errorf(i18n.T("template_plugin_error"), namespace, err)
			}
			return result, nil
		}
	}

	switch body {
	case "input", InputSentinel:
		return input, nil
	default:
		val, ok := variables[body]
		if !ok && strings.ContainsAny(body, "{}") && !strings.Contains(body, "{{") && !strings.Contains(body, "}}") {
			// A body with a single brace, as in the LaTeX text "{{a}{b}}", is
			// not a variable name. It stays as literal text. A body with "{{"
			// or "}}" gives an error, because a value can try to close a token.
			return full, nil
		}
		if !ok {
			return "", fmt.Errorf(i18n.T("template_missing_required_variable"), src)
		}
		return val, nil
	}
}
