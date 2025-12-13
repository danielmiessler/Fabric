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

// FilterByVendor returns a new VendorsModels containing only the specified vendor's models.
// Vendor matching is case-insensitive (e.g., "OpenAI", "openai", and "OPENAI" all match).
// If the vendor is not found, an empty VendorsModels is returned.
func (o *VendorsModels) FilterByVendor(vendor string) *VendorsModels {
	filtered := NewVendorsModels()
	for _, groupItems := range o.GroupsItems {
		if strings.EqualFold(groupItems.Group, vendor) {
			filtered.AddGroupItems(groupItems.Group, groupItems.Items...)
			break
		}
	}
	return filtered
}

// FindModelNameCaseInsensitive returns the actual model name from available models,
// matching case-insensitively. Returns empty string if not found.
// For example, if the available models contain "gpt-4o" and user queries "GPT-4O",
// this returns "gpt-4o" (the actual model name that should be sent to the API).
func (o *VendorsModels) FindModelNameCaseInsensitive(modelQuery string) string {
	for _, groupItems := range o.GroupsItems {
		for _, item := range groupItems.Items {
			if strings.EqualFold(item, modelQuery) {
				return item
			}
		}
	}
	return ""
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
