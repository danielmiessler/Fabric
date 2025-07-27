package tools

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
)

func NeeDefaults(getVendorsModels func() (*ai.VendorsModels, error)) (ret *Defaults) {
	vendorName := "Default"
	ret = &Defaults{
		PluginBase: &plugins.PluginBase{
			Name:             vendorName,
			SetupDescription: "Default AI Vendor and Model [required]",
			EnvNamePrefix:    plugins.BuildEnvVariablePrefix(vendorName),
		},
		GetVendorsModels: getVendorsModels,
	}

	ret.Vendor = ret.AddSetting("Vendor", true)

	ret.Model = ret.AddSetupQuestionCustom("Model", true,
		"Enter the index the name of your default model")

	ret.ModelContextLength = ret.AddSetupQuestionCustom("Model Context Length", false,
		"Enter model context length")

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

	//verify
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
