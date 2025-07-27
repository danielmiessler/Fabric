package ai

import (
	"github.com/danielmiessler/fabric/internal/util"
)

func NewVendorsModels() *VendorsModels {
	return &VendorsModels{GroupsItemsSelectorString: util.NewGroupsItemsSelectorString("Available models")}
}

type VendorsModels struct {
	*util.GroupsItemsSelectorString
}
