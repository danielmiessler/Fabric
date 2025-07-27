package template

import (
	"strings"
	"testing"
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
		// Basic variable substitution
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

		// Nested variable substitution
		{
			name:     "nested variables",
			template: "{{outer{{inner}}}}",
			vars: map[string]string{
				"inner":    "foo",    // First resolution
				"outerfoo": "result", // Second resolution
			},
			want: "result",
		},

		// Plugin operations
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

		// Multiple operations
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

		// Error cases
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

		// Edge cases
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyTemplate(tt.template, tt.vars, tt.input)

			// Check error cases
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

			// Check result
			if got != tt.want {
				t.Errorf("ApplyTemplate() = %q, want %q", got, tt.want)
			}
		})
	}
}
