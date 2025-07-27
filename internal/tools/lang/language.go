package lang

import (
	"github.com/danielmiessler/fabric/internal/plugins"
	"golang.org/x/text/language"
)

func NewLanguage() (ret *Language) {

	label := "Language"
	ret = &Language{}

	ret.PluginBase = &plugins.PluginBase{
		Name:             label,
		SetupDescription: "Language - Default AI Vendor Output Language",
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		ConfigureCustom:  ret.configure,
	}

	ret.DefaultLanguage = ret.AddSetupQuestionCustom("Output", false,
		"Enter your default output language (for example: zh_CN)")

	return
}

type Language struct {
	*plugins.PluginBase
	DefaultLanguage *plugins.SetupQuestion
}

func (o *Language) configure() error {
	if o.DefaultLanguage.Value != "" {
		langTag, err := language.Parse(o.DefaultLanguage.Value)
		if err == nil {
			o.DefaultLanguage.Value = langTag.String()
		} else {
			o.DefaultLanguage.Value = ""
		}
	}

	return nil
}
