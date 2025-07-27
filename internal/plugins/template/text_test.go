package template

import (
	"testing"
)

func TestTextPlugin(t *testing.T) {
	plugin := &TextPlugin{}

	tests := []struct {
		name      string
		operation string
		value     string
		want      string
		wantErr   bool
	}{
		// Upper tests
		{
			name:      "upper basic",
			operation: "upper",
			value:     "hello",
			want:      "HELLO",
		},
		{
			name:      "upper mixed case",
			operation: "upper",
			value:     "hElLo",
			want:      "HELLO",
		},

		// Lower tests
		{
			name:      "lower basic",
			operation: "lower",
			value:     "HELLO",
			want:      "hello",
		},
		{
			name:      "lower mixed case",
			operation: "lower",
			value:     "hElLo",
			want:      "hello",
		},

		// Title tests
		{
			name:      "title basic",
			operation: "title",
			value:     "hello world",
			want:      "Hello World",
		},
		{
			name:      "title with apostrophe",
			operation: "title",
			value:     "o'reilly's book",
			want:      "O'Reilly's Book",
		},

		// Trim tests
		{
			name:      "trim spaces",
			operation: "trim",
			value:     "  hello  ",
			want:      "hello",
		},
		{
			name:      "trim newlines",
			operation: "trim",
			value:     "\nhello\n",
			want:      "hello",
		},

		// Error cases
		{
			name:      "empty value",
			operation: "upper",
			value:     "",
			wantErr:   true,
		},
		{
			name:      "unknown operation",
			operation: "invalid",
			value:     "test",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := plugin.Apply(tt.operation, tt.value)

			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("TextPlugin.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check successful cases
			if err == nil && got != tt.want {
				t.Errorf("TextPlugin.Apply() = %q, want %q", got, tt.want)
			}
		})
	}
}
