package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	specialChars  = map[rune]bool{':': true, '"': true, '\\': true, '\n': true, '\t': true, '\r': true, '[': true, ']': true, '{': true, '}': true}
	numericRe     = regexp.MustCompile(`^-?\d+(?:\.\d+)?(?:[eE][+-]?\d+)?$`)
	leadingZeroRe = regexp.MustCompile(`^0\d+$`)
)

func needsQuoting(value, delimiter string) bool {
	if value == "" || value == "true" || value == "false" || value == "null" {
		return true
	}
	if value[0] == ' ' || value[len(value)-1] == ' ' || value[0] == '-' {
		return true
	}
	for _, c := range value {
		if specialChars[c] {
			return true
		}
	}
	return strings.Contains(value, delimiter) || numericRe.MatchString(value) || leadingZeroRe.MatchString(value)
}

func escape(value string) string {
	r := strings.NewReplacer("\\", "\\\\", "\"", "\\\"", "\n", "\\n", "\r", "\\r", "\t", "\\t")
	return r.Replace(value)
}

func quote(value, delimiter string) string {
	if needsQuoting(value, delimiter) {
		return fmt.Sprintf(`"%s"`, escape(value))
	}
	return value
}

func encodeValue(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case bool:
		if v {
			return "true"
		}
		return "false"
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%g", v)
	case string:
		return quote(v, ",")
	default:
		return quote(fmt.Sprintf("%v", v), ",")
	}
}

func isPrimitive(value any) bool {
	switch value.(type) {
	case nil, bool, int, int64, float64, string:
		return true
	}
	return false
}

func isTabular(arr []any) bool {
	if len(arr) == 0 {
		return false
	}
	first, ok := arr[0].(map[string]any)
	if !ok {
		return false
	}
	keys := make(map[string]bool)
	for k := range first {
		keys[k] = true
	}
	for _, item := range arr {
		m, ok := item.(map[string]any)
		if !ok || len(m) != len(keys) {
			return false
		}
		for k, v := range m {
			if !keys[k] || !isPrimitive(v) {
				return false
			}
		}
	}
	return true
}

func keys(m map[string]any) []string {
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func encodeList(arr []any, indent int) string {
	if len(arr) == 0 {
		return "[0]:"
	}

	allPrimitive := true
	for _, item := range arr {
		if !isPrimitive(item) {
			allPrimitive = false
			break
		}
	}
	if allPrimitive {
		vals := make([]string, len(arr))
		for i, v := range arr {
			vals[i] = encodeValue(v)
		}
		return fmt.Sprintf("[%d]: %s", len(arr), strings.Join(vals, ","))
	}

	if isTabular(arr) {
		first := arr[0].(map[string]any)
		ks := keys(first)
		var rows []string
		for _, item := range arr {
			m := item.(map[string]any)
			vals := make([]string, len(ks))
			for i, k := range ks {
				vals[i] = encodeValue(m[k])
			}
			rows = append(rows, strings.Join(vals, ","))
		}
		return fmt.Sprintf("[%d]{%s}:\n%s", len(arr), strings.Join(ks, ","), strings.Join(rows, "\n"))
	}

	pad := strings.Repeat("  ", indent)
	lines := []string{fmt.Sprintf("[%d]:", len(arr))}
	for _, item := range arr {
		if m, ok := item.(map[string]any); ok {
			nested := strings.Split(encodeDict(m, indent+1), "\n")
			lines = append(lines, fmt.Sprintf("%s  %s", pad, nested[0]))
			for _, ln := range nested[1:] {
				lines = append(lines, fmt.Sprintf("%s    %s", pad, ln))
			}
		} else {
			lines = append(lines, fmt.Sprintf("%s  %s", pad, encodeValue(item)))
		}
	}
	return strings.Join(lines, "\n")
}

func encodeDict(obj map[string]any, indent int) string {
	if len(obj) == 0 {
		return ""
	}
	var lines []string
	for _, key := range keys(obj) {
		value := obj[key]
		switch v := value.(type) {
		case map[string]any:
			lines = append(lines, key+":")
			for _, ln := range strings.Split(encodeDict(v, indent+1), "\n") {
				if ln != "" {
					lines = append(lines, "  "+ln)
				}
			}
		case []any:
			encoded := encodeList(v, indent+1)
			if strings.Contains(encoded, "\n") {
				parts := strings.Split(encoded, "\n")
				lines = append(lines, key+parts[0])
				for _, ln := range parts[1:] {
					lines = append(lines, "  "+ln)
				}
			} else {
				lines = append(lines, key+encoded)
			}
		default:
			lines = append(lines, fmt.Sprintf("%s: %s", key, encodeValue(v)))
		}
	}
	return strings.Join(lines, "\n")
}

func PromptToTOON(p *FabricPrompt) string {
	var lines []string

	lines = append(lines, fmt.Sprintf("role: %s", encodeValue(p.Role)))

	if len(p.Expertise) > 0 {
		vals := make([]string, len(p.Expertise))
		for i, e := range p.Expertise {
			vals[i] = encodeValue(e)
		}
		lines = append(lines, fmt.Sprintf("expertise[%d]: %s", len(p.Expertise), strings.Join(vals, ",")))
	} else {
		lines = append(lines, "expertise[0]:")
	}

	lines = append(lines, fmt.Sprintf("purpose: %s", encodeValue(p.Purpose)))

	if len(p.Steps) > 0 {
		ks := keys(p.Steps[0])
		lines = append(lines, fmt.Sprintf("steps[%d]{%s}:", len(p.Steps), strings.Join(ks, ",")))
		for _, step := range p.Steps {
			vals := make([]string, len(ks))
			for i, k := range ks {
				vals[i] = encodeValue(step[k])
			}
			lines = append(lines, "  "+strings.Join(vals, ","))
		}
	} else {
		lines = append(lines, "steps[0]:")
	}

	lines = append(lines, fmt.Sprintf("output_format: %s", p.OutputFormat))

	if len(p.OutputSections) > 0 {
		lines = append(lines, fmt.Sprintf("output_sections[%d]:", len(p.OutputSections)))
	} else {
		lines = append(lines, "output_sections[0]:")
	}

	if len(p.OutputInstructions) > 0 {
		ks := keys(p.OutputInstructions[0])
		lines = append(lines, fmt.Sprintf("output_instructions[%d]{%s}:", len(p.OutputInstructions), strings.Join(ks, ",")))
		for _, inst := range p.OutputInstructions {
			vals := make([]string, len(ks))
			for i, k := range ks {
				vals[i] = encodeValue(inst[k])
			}
			lines = append(lines, "  "+strings.Join(vals, ","))
		}
	} else {
		lines = append(lines, "output_instructions[0]:")
	}

	if len(p.Restrictions) > 0 {
		ks := keys(p.Restrictions[0])
		lines = append(lines, fmt.Sprintf("restrictions[%d]{%s}:", len(p.Restrictions), strings.Join(ks, ",")))
		for _, r := range p.Restrictions {
			vals := make([]string, len(ks))
			for i, k := range ks {
				vals[i] = encodeValue(r[k])
			}
			lines = append(lines, "  "+strings.Join(vals, ","))
		}
	} else {
		lines = append(lines, "restrictions[0]:")
	}

	if p.ThinkingInstruction != "" {
		lines = append(lines, fmt.Sprintf("thinking_instruction: %s", encodeValue(p.ThinkingInstruction)))
	}

	return strings.Join(lines, "\n")
}
