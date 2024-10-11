package lang

import (
	"github.com/danielmiessler/fabric/common"
	"golang.org/x/text/language"
)

func NewLanguage() (ret *Language) {

	label := "Language"
	ret = &Language{}

	ret.Configurable = &common.Configurable{
		Label:           label,
		EnvNamePrefix:   common.BuildEnvVariablePrefix(label),
		ConfigureCustom: ret.configure,
	}

	ret.DefaultLanguage = ret.Configurable.AddSetupQuestionCustom("Output", false,
		"Enter your default want output lang (for example: zh_CN)")

	return
}

type Language struct {
	*common.Configurable
	DefaultLanguage *common.SetupQuestion
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
