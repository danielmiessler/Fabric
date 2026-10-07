package template

import (
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// TestSinglePassValueNotReexpanded checks that a plugin or extension token in
// a variable value stays as literal text, and that the plugin does not run.
// A REST client can set the variables. Thus the engine must not resolve the
// tokens in a variable value.
func TestSinglePassValueNotReexpanded(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "ENV_VALUE_FROM_PLUGIN")

	cases := []struct {
		name     string
		template string
		vars     map[string]string
		literal  string
	}{
		{
			name:     "sys env in variable value",
			template: "{{value}}",
			vars:     map[string]string{"value": "{{plugin:sys:env:FABRIC_TEST_VALUE}}"},
			literal:  "{{plugin:sys:env:FABRIC_TEST_VALUE}}",
		},
		{
			name:     "file read in variable value",
			template: "x {{value}} y",
			vars:     map[string]string{"value": "{{plugin:file:read:notes.txt}}"},
			literal:  "{{plugin:file:read:notes.txt}}",
		},
		{
			name:     "ext in variable value",
			template: "{{value}}",
			vars:     map[string]string{"value": "{{ext:anything:run:x}}"},
			literal:  "{{ext:anything:run:x}}",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ApplyTemplate(tc.template, tc.vars, "")
			if err != nil {
				t.Fatalf("ApplyTemplate() unexpected error = %v", err)
			}
			if !strings.Contains(got, tc.literal) {
				t.Errorf("value was expanded again: got %q, want it to contain the literal token %q", got, tc.literal)
			}
			if strings.Contains(got, "ENV_VALUE_FROM_PLUGIN") {
				t.Errorf("plugin ran on a variable value: got %q", got)
			}
		})
	}
}

// TestSinglePassValueCannotCloseToken checks that a variable value in a
// nested token cannot close the outer token and start a second plugin call.
// For the first value, matchTriple must match the full token. For the second
// value, the engine must not scan the value again. The loop before the single
// pass gave the environment value for the second value.
func TestSinglePassValueCannotCloseToken(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "ENV_VALUE_FROM_PLUGIN")

	// The template has one nested plugin call. Each value tries to close that
	// call and start a sys plugin call.
	for _, value := range []string{
		"}}{{plugin:sys:env:FABRIC_TEST_VALUE",
		"x}}{{plugin:sys:env:FABRIC_TEST_VALUE",
	} {
		got, err := ApplyTemplate("{{plugin:text:upper:{{name}}}}", map[string]string{"name": value}, "")
		// An error or other output is correct. The environment value is not.
		if err == nil && strings.Contains(got, "ENV_VALUE_FROM_PLUGIN") {
			t.Errorf("the sys plugin ran from the value %q: got %q", value, got)
		}
	}
}

// TestSinglePassAuthoredTokenStillRuns checks that a plugin token in the
// template text still runs. The single pass only stops the scan of values.
func TestSinglePassAuthoredTokenStillRuns(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "ENV_VALUE_FROM_PLUGIN")

	got, err := ApplyTemplate("env={{plugin:sys:env:FABRIC_TEST_VALUE}}", map[string]string{}, "")
	if err != nil {
		t.Fatalf("ApplyTemplate() unexpected error = %v", err)
	}
	if got != "env=ENV_VALUE_FROM_PLUGIN" {
		t.Errorf("authored plugin token did not run: got %q, want %q", got, "env=ENV_VALUE_FROM_PLUGIN")
	}
}

// TestSinglePassAuthoredNestedVariable checks that nested tokens in the
// template text still work. The inner variable resolves first, and then the
// outer plugin runs on the result.
func TestSinglePassAuthoredNestedVariable(t *testing.T) {
	got, err := ApplyTemplate("{{plugin:text:upper:{{name}}}}", map[string]string{"name": "world"}, "")
	if err != nil {
		t.Fatalf("ApplyTemplate() unexpected error = %v", err)
	}
	if got != "WORLD" {
		t.Errorf("authored nested token broke: got %q, want %q", got, "WORLD")
	}
}

// TestNoSystemPlugins checks that ApplyTemplateNoSystemPlugins refuses each
// plugin that can read the host or the network, and each extension. The
// text and datetime plugins still run.
func TestNoSystemPlugins(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "secret")

	for _, tmpl := range []string{
		"{{plugin:sys:env:FABRIC_TEST_VALUE}}",
		"{{plugin:file:read:notes.txt}}",
		"{{plugin:fetch:get:https://example.com/}}",
		"{{ext:anything:run:x}}",
	} {
		got, err := ApplyTemplateNoSystemPlugins("a "+tmpl, nil, "")
		if err == nil || strings.Contains(got, "secret") {
			t.Errorf("%s: got %q, %v; want an error", tmpl, got, err)
		}
	}

	got, err := ApplyTemplateNoSystemPlugins("{{plugin:text:upper:{{name}}}}", map[string]string{"name": "ok"}, "")
	if err != nil || got != "OK" {
		t.Errorf("text plugin: got %q, %v; want %q", got, err, "OK")
	}
}

// TestEngineDispatchesSystemPlugins checks that a file, fetch or datetime
// token in a template gets to its plugin through ApplyTemplate, and that the
// plugin checks still apply.
func TestEngineDispatchesSystemPlugins(t *testing.T) {
	server := newTextServer("loopback content")
	defer server.Close()

	cases := []struct {
		name     string
		template string
		wantErr  string // "" for success
	}{
		{"file read out of root", "{{plugin:file:read:../secret.txt}}", i18n.T("template_file_error_path_contains_parent_ref")},
		{"fetch of loopback address", "{{plugin:fetch:get:" + server.URL + "}}", i18n.T("util_error_non_public_address")},
		{"datetime", "{{plugin:datetime:unix}}", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ApplyTemplate(tc.template, nil, "")
			if tc.wantErr == "" {
				if err != nil || got == "" || strings.Contains(got, "{{") {
					t.Fatalf("ApplyTemplate() = %q, %v; want a rendered value", got, err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("ApplyTemplate() error = %v, want %q", err, tc.wantErr)
			}
		})
	}
}

// TestSinglePassValueAsFullBody checks that a token with only an inner token
// as its body does not resolve the inner value as a token.
func TestSinglePassValueAsFullBody(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "ENV_VALUE_FROM_PLUGIN")

	cases := []struct {
		vars map[string]string
		want string
	}{
		{map[string]string{"name": "plugin:sys:env:FABRIC_TEST_VALUE"}, "{{plugin:sys:env:FABRIC_TEST_VALUE}}"},
		{map[string]string{"name": "other", "other": "SECOND"}, "{{other}}"},
	}
	for _, tc := range cases {
		got, err := ApplyTemplate("{{{{name}}}}", tc.vars, "")
		if err != nil || got != tc.want {
			t.Errorf("ApplyTemplate(%v) = %q, %v; want %q", tc.vars, got, err, tc.want)
		}
	}
}

// TestMissingVariableErrorHidesInnerValue checks that the error for a missing
// variable shows the template text, not the values of the inner tokens.
func TestMissingVariableErrorHidesInnerValue(t *testing.T) {
	t.Setenv("FABRIC_TEST_VALUE", "ENV_VALUE_FROM_PLUGIN")

	_, err := ApplyTemplate("{{x {{plugin:sys:env:FABRIC_TEST_VALUE}}}}", nil, "")
	if err == nil || strings.Contains(err.Error(), "ENV_VALUE_FROM_PLUGIN") {
		t.Errorf("ApplyTemplate() error = %v; want an error without the value", err)
	}
}
