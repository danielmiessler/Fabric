package plugins

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const AnswerReset = "reset"
const SettingTypeBool = "bool"

type Plugin interface {
	GetName() string
	GetSetupDescription() string
	IsConfigured() bool
	Configure() error
	Setup() error
	SetupFillEnvFileContent(*bytes.Buffer)
}

type PluginBase struct {
	Settings
	SetupQuestions

	Name             string
	SetupDescription string
	EnvNamePrefix    string

	ConfigureCustom func() error
}

func (o *PluginBase) GetName() string {
	return o.Name
}

func (o *PluginBase) GetSetupDescription() (ret string) {
	if ret = o.SetupDescription; ret == "" {
		ret = o.GetName()
	}
	return
}

func (o *PluginBase) AddSetting(name string, required bool) (ret *Setting) {
	ret = NewSetting(fmt.Sprintf("%v%v", o.EnvNamePrefix, BuildEnvVariable(name)), required)
	o.Settings = append(o.Settings, ret)
	return
}

func (o *PluginBase) AddSetupQuestion(name string, required bool) (ret *SetupQuestion) {
	return o.AddSetupQuestionCustom(name, required, "")
}

func (o *PluginBase) AddSetupQuestionCustom(name string, required bool, question string) (ret *SetupQuestion) {
	setting := o.AddSetting(name, required)
	ret = &SetupQuestion{Setting: setting, Question: question}
	if ret.Question == "" {
		ret.Question = fmt.Sprintf("Enter your %v %v", o.Name, strings.ToUpper(name))
	}
	o.SetupQuestions = append(o.SetupQuestions, ret)
	return
}

func (o *PluginBase) AddSetupQuestionBool(name string, required bool) (ret *SetupQuestion) {
	return o.AddSetupQuestionCustomBool(name, required, "")
}

func (o *PluginBase) AddSetupQuestionCustomBool(name string, required bool, question string) (ret *SetupQuestion) {
	setting := o.AddSetting(name, required)
	setting.Type = SettingTypeBool
	ret = &SetupQuestion{Setting: setting, Question: question}
	if ret.Question == "" {
		ret.Question = fmt.Sprintf("Enable %v %v (true/false)", o.Name, strings.ToUpper(name))
	}
	o.SetupQuestions = append(o.SetupQuestions, ret)
	return
}

func (o *PluginBase) Configure() (err error) {
	if err = o.Settings.Configure(); err != nil {
		return
	}

	if o.ConfigureCustom != nil {
		err = o.ConfigureCustom()
	}
	return
}

func (o *PluginBase) Setup() (err error) {
	if err = o.Ask(o.Name); err != nil {
		return
	}

	err = o.Configure()
	return
}

func (o *PluginBase) SetupOrSkip() (err error) {
	if err = o.Setup(); err != nil {
		fmt.Printf("[%v] skipped\n", o.GetName())
	}
	return
}

func (o *PluginBase) SetupFillEnvFileContent(fileEnvFileContent *bytes.Buffer) {
	o.Settings.FillEnvFileContent(fileEnvFileContent)
}

func NewSetting(envVariable string, required bool) *Setting {
	return &Setting{
		EnvVariable: envVariable,
		Required:    required,
	}
}

// In plugins/plugin.go

type Setting struct {
	EnvVariable string
	Value       string
	Required    bool
	Type        string // "string" (default), "bool"
}

func (o *Setting) IsValid() bool {
	if o.Type == SettingTypeBool {
		_, err := ParseBool(o.Value)
		return (err == nil) || !o.Required
	}
	return o.IsDefined() || !o.Required
}

func (o *Setting) Print() {
	if o.Type == SettingTypeBool {
		v, _ := ParseBool(o.Value)
		fmt.Printf("%v: %v\n", o.EnvVariable, v)
	} else {
		fmt.Printf("%v: %v\n", o.EnvVariable, o.Value)
	}
}

func (o *Setting) FillEnvFileContent(buffer *bytes.Buffer) {
	if o.IsDefined() {
		buffer.WriteString(o.EnvVariable)
		buffer.WriteString("=")
		if o.Type == SettingTypeBool {
			v, _ := ParseBool(o.Value)
			buffer.WriteString(fmt.Sprintf("%v", v))
		} else {
			buffer.WriteString(o.Value)
		}
		buffer.WriteString("\n")
	}
}

func ParseBoolElseFalse(val string) (ret bool) {
	ret, _ = ParseBool(val)
	return
}

func ParseBool(val string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(val)) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	}
	return false, fmt.Errorf("invalid bool: %q", val)
}

type SetupQuestion struct {
	*Setting
	Question string
}

func (o *SetupQuestion) Ask(label string) (err error) {
	var prefix string
	if label != "" {
		prefix = fmt.Sprintf("[%v] ", label)
	} else {
		prefix = ""
	}
	fmt.Println()
	if o.Type == SettingTypeBool {
		current := "false"
		if v, err := ParseBool(o.Value); err == nil && v {
			current = "true"
		}
		fmt.Printf("%v%v (true/false, leave empty for '%s' or type '%v' to remove the value):\n",
			prefix, o.Question, current, AnswerReset)
	} else if o.Value != "" {
		fmt.Printf("%v%v (leave empty for '%s' or type '%v' to remove the value):\n",
			prefix, o.Question, o.Value, AnswerReset)
	} else {
		fmt.Printf("%v%v (leave empty to skip):\n", prefix, o.Question)
	}
	var answer string
	fmt.Scanln(&answer)
	answer = strings.TrimRight(answer, "\n")
	if answer == "" {
		answer = o.Value
	} else if strings.ToLower(answer) == AnswerReset {
		answer = ""
	}
	err = o.OnAnswer(answer)
	return
}

func (o *SetupQuestion) OnAnswer(answer string) (err error) {
	if o.Type == SettingTypeBool {
		if answer == "" {
			o.Value = ""
		} else {
			_, err := ParseBool(answer)
			if err != nil {
				return fmt.Errorf("invalid boolean value: %v", answer)
			}
			o.Value = strings.ToLower(answer)
		}
	} else {
		o.Value = answer
	}
	if o.EnvVariable != "" {
		if err = os.Setenv(o.EnvVariable, o.Value); err != nil {
			return
		}
	}
	err = o.IsValidErr()
	return
}

func (o *Setting) IsValidErr() (err error) {
	if !o.IsValid() {
		err = fmt.Errorf("%v=%v, is not valid", o.EnvVariable, o.Value)
	}
	return
}

func (o *Setting) IsDefined() bool {
	return o.Value != ""
}

func (o *Setting) Configure() error {
	envValue := os.Getenv(o.EnvVariable)
	if envValue != "" {
		o.Value = envValue
	}
	return o.IsValidErr()
}

func NewSetupQuestion(question string) *SetupQuestion {
	return &SetupQuestion{Setting: &Setting{}, Question: question}
}

type Settings []*Setting

func (o Settings) IsConfigured() (ret bool) {
	ret = true
	for _, setting := range o {
		if ret = setting.IsValid(); !ret {
			break
		}
	}
	return
}

func (o Settings) Configure() (err error) {
	for _, setting := range o {
		if err = setting.Configure(); err != nil {
			break
		}
	}
	return
}

func (o Settings) FillEnvFileContent(buffer *bytes.Buffer) {
	for _, setting := range o {
		setting.FillEnvFileContent(buffer)
	}
}

type SetupQuestions []*SetupQuestion

func (o SetupQuestions) Ask(label string) (err error) {
	fmt.Println()
	fmt.Printf("[%v]\n", label)
	for _, question := range o {
		if err = question.Ask(""); err != nil {
			break
		}
	}
	return
}

func BuildEnvVariablePrefix(name string) (ret string) {
	ret = BuildEnvVariable(name)
	if ret != "" {
		ret += "_"
	}
	return
}

func BuildEnvVariable(name string) string {
	name = strings.TrimSpace(name)
	return strings.ReplaceAll(strings.ToUpper(name), " ", "_")
}
