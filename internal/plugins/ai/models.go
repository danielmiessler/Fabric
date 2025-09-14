package ai

import (
	"fmt"
	"sort"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/util"
)

func NewVendorsModels() *VendorsModels {
	return &VendorsModels{GroupsItemsSelectorString: util.NewGroupsItemsSelectorString(i18n.T("available_models_header"))}
}

type VendorsModels struct {
	*util.GroupsItemsSelectorString
}

// PrintWithVendor prints models including their vendor on each line.
// When shellCompleteList is true, output is suitable for shell completion.
// Default vendor and model are highlighted with an asterisk.
func (o *VendorsModels) PrintWithVendor(shellCompleteList bool, defaultVendor, defaultModel string) {
	if !shellCompleteList {
		fmt.Printf("%s:\n\n", o.SelectionLabel)
	}

	var currentItemIndex int

	sortedGroups := make([]*util.GroupItems[string], len(o.GroupsItems))
	copy(sortedGroups, o.GroupsItems)
	sort.SliceStable(sortedGroups, func(i, j int) bool {
		return strings.ToLower(sortedGroups[i].Group) < strings.ToLower(sortedGroups[j].Group)
	})

	for _, groupItems := range sortedGroups {
		items := make([]string, len(groupItems.Items))
		copy(items, groupItems.Items)
		sort.SliceStable(items, func(i, j int) bool {
			return strings.ToLower(items[i]) < strings.ToLower(items[j])
		})
		for _, item := range items {
			currentItemIndex++
			if shellCompleteList {
				fmt.Printf("%s|%s\n", groupItems.Group, item)
			} else {
				mark := "       "
				if strings.EqualFold(groupItems.Group, defaultVendor) && strings.EqualFold(item, defaultModel) {
					mark = "      *"
				}
				fmt.Printf("%s\t[%d]\t%s|%s\n", mark, currentItemIndex, groupItems.Group, item)
			}
		}
	}
}
