package tools

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
)

func NeeDefaults(getVendorsModels func() (*ai.VendorsModels, error)) (ret *Defaults) {
	vendorName := "Default"
	ret = &Defaults{
		PluginBase: &plugins.PluginBase{
			Name:             vendorName,
			SetupDescription: i18n.T("defaults_setup_description") + " " + i18n.T("required_marker"),
			EnvNamePrefix:    plugins.BuildEnvVariablePrefix(vendorName),
		},
		GetVendorsModels: getVendorsModels,
	}

	ret.Vendor = ret.AddSetting("Vendor", true)

	ret.Model = ret.AddSetupQuestionWithEnvName("Model", true,
		i18n.T("defaults_model_question"))

	ret.ModelContextLength = ret.AddSetupQuestionWithEnvName("Model Context Length", false,
		i18n.T("defaults_model_context_length_question"))

	return
}

type Defaults struct {
	*plugins.PluginBase

	Vendor             *plugins.Setting
	Model              *plugins.SetupQuestion
	ModelContextLength *plugins.SetupQuestion
	GetVendorsModels   func() (*ai.VendorsModels, error)
}

func (o *Defaults) Setup() (err error) {
	var vendorsModels *ai.VendorsModels
	if vendorsModels, err = o.GetVendorsModels(); err != nil {
		return
	}

	vendorsModels.Print(false)

	if err = o.Ask(o.Name); err != nil {
		return
	}

	index, parseErr := strconv.Atoi(o.Model.Value)
	if parseErr == nil {
		if o.Vendor.Value, o.Model.Value, err = vendorsModels.GetGroupAndItemByItemNumber(index); err != nil {
			return
		}
	} else {
		o.Vendor.Value = vendorsModels.FindGroupsByItemFirst(o.Model.Value)
	}

	vendorNames := vendorsModels.FindGroupsByItem(o.Model.Value)
	if len(vendorNames) == 0 {
		err = errors.Errorf("You need to chose an available default model.")
		return
	}

	fmt.Println()
	o.Vendor.Print()
	o.Model.Print()

	return
}
