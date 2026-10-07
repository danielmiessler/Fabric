package template

import (
	"strings"
	"testing"
	"time"
)

func TestApplyTemplate(t *testing.T) {
	tests := []struct {
		name        string
		template    string
		vars        map[string]string
		input       string
		want        string
		wantErr     bool
		errContains string
	}{
		{
			name:     "simple variable",
			template: "Hello {{name}}!",
			vars:     map[string]string{"name": "World"},
			want:     "Hello World!",
		},
		{
			name:     "multiple variables",
			template: "{{greeting}} {{name}}!",
			vars: map[string]string{
				"greeting": "Hello",
				"name":     "World",
			},
			want: "Hello World!",
		},
		{
			name:     "special input variable",
			template: "Content: {{input}}",
			input:    "test content",
			want:     "Content: test content",
		},

		{
			name:     "nested variables",
			template: "{{outer{{inner}}}}",
			vars: map[string]string{
				"inner":    "foo",    // First resolution
				"outerfoo": "result", // Second resolution
			},
			want: "result",
		},

		{
			name:     "simple text plugin",
			template: "{{plugin:text:upper:hello}}",
			want:     "HELLO",
		},
		{
			name:     "text plugin with variable",
			template: "{{plugin:text:upper:{{name}}}}",
			vars:     map[string]string{"name": "world"},
			want:     "WORLD",
		},
		{
			name:     "plugin with dynamic operation",
			template: "{{plugin:text:{{operation}}:hello}}",
			vars:     map[string]string{"operation": "upper"},
			want:     "HELLO",
		},

		{
			name:     "multiple plugins",
			template: "A:{{plugin:text:upper:hello}} B:{{plugin:text:lower:WORLD}}",
			want:     "A:HELLO B:world",
		},
		{
			name:     "nested plugins",
			template: "{{plugin:text:upper:{{plugin:text:lower:HELLO}}}}",
			want:     "HELLO",
		},

		{
			name:        "missing variable",
			template:    "Hello {{name}}!",
			wantErr:     true,
			errContains: "missing required variable",
		},
		{
			name:        "unknown plugin",
			template:    "{{plugin:invalid:op:value}}",
			wantErr:     true,
			errContains: "unknown plugin namespace",
		},
		{
			name:        "unknown plugin operation",
			template:    "{{plugin:text:invalid:value}}",
			wantErr:     true,
			errContains: "unknown text operation",
		},
		{
			name:        "nested plugin error",
			template:    "{{plugin:text:upper:{{plugin:invalid:op:value}}}}",
			wantErr:     true,
			errContains: "unknown plugin namespace",
		},

		{
			name:     "empty template",
			template: "",
			want:     "",
		},
		{
			name:     "no substitutions needed",
			template: "plain text",
			want:     "plain text",
		},

		{
			name:     "open braces with no close braces",
			template: "x{{{{a}}",
			vars:     map[string]string{"a": "V"},
			want:     "x{{V",
		},
		{
			name:     "close braces with no open braces",
			template: "{{a}}}}",
			vars:     map[string]string{"a": "V"},
			want:     "V}}",
		},
		{
			name:     "triple braces",
			template: "{{{a}}}",
			vars:     map[string]string{"a": "V"},
			want:     "{V}",
		},
		{
			name:     "triple open braces",
			template: "{{{a}}",
			vars:     map[string]string{"a": "V"},
			want:     "{V",
		},
		{
			name:     "single braces in token",
			template: "{{a}{b}}",
			want:     "{{a}{b}}",
		},
		{
			name:     "value is not a variable name",
			template: "{{{{inner}}}}",
			vars:     map[string]string{"inner": "a", "a": "V"},
			want:     "{{a}}",
		},
		{
			name:     "nested variable name with template text",
			template: "{{x_{{inner}}}}",
			vars:     map[string]string{"inner": "a", "x_a": "V"},
			want:     "V",
		},
		{
			name:     "empty token",
			template: "{{}}",
			want:     "{{}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyTemplate(tt.template, tt.vars, tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errContains != "" {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error %q should contain %q", err.Error(), tt.errContains)
				}
				return
			}

			if got != tt.want {
				t.Errorf("ApplyTemplate() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestApplyTemplateLongBraceRun checks that the time of the scan increases
// linearly with the length of the content. ApplyTemplate must give 1 MiB of
// "{" back as literal text in less than 5 seconds.
func TestApplyTemplateLongBraceRun(t *testing.T) {
	content := strings.Repeat("{", 1<<20)
	if out, err := applyWithin(t, content, nil); err != nil || out != content {
		t.Errorf("ApplyTemplate() changed the text, error = %v", err)
	}
}

// TestCyclicVariablesStayLiteral checks variable values that contain a
// variable token, for example a={{a}}. The engine does not scan a value
// again. Thus ApplyTemplate stops, and the value stays as literal text.
func TestCyclicVariablesStayLiteral(t *testing.T) {
	tests := []struct {
		name string
		vars map[string]string
		want string
	}{
		{"self reference", map[string]string{"a": "{{a}}"}, "{{a}}"},
		{"mutual cycle", map[string]string{"a": "{{b}}", "b": "{{a}}"}, "{{b}}"},
		{"expanding self reference", map[string]string{"a": "{{a}}{{a}}"}, "{{a}}{{a}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyWithin(t, "{{a}}", tt.vars)
			if err != nil || got != tt.want {
				t.Errorf("ApplyTemplate() = %q, %v, want %q", got, err, tt.want)
			}
		})
	}
}

// TestApplyTemplateInput checks that ApplyTemplateInput replaces variables,
// but gives plugin and extension tokens back as literal text. This is also
// true for a token that the text of an inner token makes.
func TestApplyTemplateInput(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "ENV_VALUE_FROM_PLUGIN")
	vars := map[string]string{"name": "world", "empty": ""}
	tests := []struct{ name, input, want string }{
		{"variable", "Hello {{name}}", "Hello world"},
		{"plugin token", "{{plugin:sys:env:FABRIC_TEST_VALUE}}", "{{plugin:sys:env:FABRIC_TEST_VALUE}}"},
		{"extension token", "{{ext:anything:run:x}}", "{{ext:anything:run:x}}"},
		{"nested plugin tokens", "{{plugin:text:upper:{{plugin:sys:env:FABRIC_TEST_VALUE}}}}", "{{plugin:text:upper:{{plugin:sys:env:FABRIC_TEST_VALUE}}}}"},
		{"variable in plugin token", "{{plugin:text:upper:{{name}}}}", "{{plugin:text:upper:world}}"},
		{"token made with input", "{{{{input}}plugin:sys:env:FABRIC_TEST_VALUE}}", "{{plugin:sys:env:FABRIC_TEST_VALUE}}"},
		{"token made with empty variable", "{{{{empty}}plugin:sys:env:FABRIC_TEST_VALUE}}", "{{plugin:sys:env:FABRIC_TEST_VALUE}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyTemplateInput(tt.input, vars)
			if err != nil || got != tt.want {
				t.Errorf("ApplyTemplateInput() = %q, %v, want %q", got, err, tt.want)
			}
		})
	}
	// An unknown variable in the input gives an error.
	if _, err := ApplyTemplateInput("{{unknown}}", vars); err == nil || !strings.Contains(err.Error(), "missing required variable") {
		t.Errorf("ApplyTemplateInput() unknown variable: error = %v, want missing required variable", err)
	}
}

// applyWithin starts ApplyTemplate with no input in a goroutine. If
// ApplyTemplate does not stop in 5 seconds, applyWithin stops the test with
// an error.
func applyWithin(t *testing.T, content string, vars map[string]string) (string, error) {
	t.Helper()
	type result struct {
		out string
		err error
	}
	done := make(chan result, 1)
	go func() {
		out, err := ApplyTemplate(content, vars, "")
		done <- result{out, err}
	}()
	select {
	case r := <-done:
		return r.out, r.err
	case <-time.After(5 * time.Second):
		t.Fatal("ApplyTemplate() did not stop in 5 seconds")
		return "", nil
	}
}
