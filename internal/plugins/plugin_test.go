package plugins

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurable_AddSetting(t *testing.T) {
	conf := &PluginBase{
		Settings:      Settings{},
		Name:          "TestConfigurable",
		EnvNamePrefix: "TEST_",
	}

	setting := conf.AddSetting("test_setting", true)
	assert.Equal(t, "TEST_TEST_SETTING", setting.EnvVariable)
	assert.True(t, setting.Required)
	assert.Contains(t, conf.Settings, setting)
}

func TestConfigurable_Configure(t *testing.T) {
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Required:    true,
	}
	conf := &PluginBase{
		Settings: Settings{setting},
		Name:     "TestConfigurable",
	}

	_ = os.Setenv("TEST_SETTING", "test_value")
	err := conf.Configure()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", setting.Value)
}

func TestConfigurable_Setup(t *testing.T) {
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Required:    false,
	}
	conf := &PluginBase{
		Settings: Settings{setting},
		Name:     "TestConfigurable",
	}

	err := conf.Setup()
	assert.NoError(t, err)
}

func TestSetting_IsValid(t *testing.T) {
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Value:       "some_value",
		Required:    true,
	}

	assert.True(t, setting.IsValid())

	setting.Value = ""
	assert.False(t, setting.IsValid())
}

func TestSetting_Configure(t *testing.T) {
	_ = os.Setenv("TEST_SETTING", "test_value")
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Required:    true,
	}
	err := setting.Configure()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", setting.Value)
}

func TestSetting_FillEnvFileContent(t *testing.T) {
	buffer := &bytes.Buffer{}
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Value:       "test_value",
	}
	setting.FillEnvFileContent(buffer)

	expected := "TEST_SETTING=test_value\n"
	assert.Equal(t, expected, buffer.String())
}

func TestSetting_Print(t *testing.T) {
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Value:       "test_value",
	}
	expected := "TEST_SETTING: test_value\n"
	fmtOutput := captureOutput(func() {
		setting.Print()
	})
	assert.Equal(t, expected, fmtOutput)
}

func TestSetupQuestion_Ask(t *testing.T) {
	setting := &Setting{
		EnvVariable: "TEST_SETTING",
		Required:    true,
	}
	question := &SetupQuestion{
		Setting:  setting,
		Question: "Enter test setting:",
	}
	input := "user_value\n"
	fmtInput := captureInput(input)
	defer fmtInput()
	err := question.Ask("TestConfigurable")
	assert.NoError(t, err)
	assert.Equal(t, "user_value", setting.Value)
}

func TestSettings_IsConfigured(t *testing.T) {
	settings := Settings{
		{EnvVariable: "TEST_SETTING1", Value: "value1", Required: true},
		{EnvVariable: "TEST_SETTING2", Value: "", Required: false},
	}

	assert.True(t, settings.IsConfigured())

	settings[0].Value = ""
	assert.False(t, settings.IsConfigured())
}

func TestSettings_Configure(t *testing.T) {
	_ = os.Setenv("TEST_SETTING", "test_value")
	settings := Settings{
		{EnvVariable: "TEST_SETTING", Required: true},
	}

	err := settings.Configure()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", settings[0].Value)
}

func TestSettings_FillEnvFileContent(t *testing.T) {
	buffer := &bytes.Buffer{}
	settings := Settings{
		{EnvVariable: "TEST_SETTING", Value: "test_value"},
	}
	settings.FillEnvFileContent(buffer)

	expected := "TEST_SETTING=test_value\n"
	assert.Equal(t, expected, buffer.String())
}

// captureOutput captures the output of a function call
func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	_ = w.Close()
	os.Stdout = stdout
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// captureInput captures the input for a function call
func captureInput(input string) func() {
	r, w, _ := os.Pipe()
	_, _ = w.WriteString(input)
	_ = w.Close()
	stdin := os.Stdin
	os.Stdin = r
	return func() {
		os.Stdin = stdin
	}
}
