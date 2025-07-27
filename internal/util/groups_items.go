package util

import (
	"fmt"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func NewGroupsItemsSelector[I any](selectionLabel string,
	getItemLabel func(I) string) *GroupsItemsSelector[I] {

	return &GroupsItemsSelector[I]{SelectionLabel: selectionLabel,
		GetItemKey:  getItemLabel,
		GroupsItems: make([]*GroupItems[I], 0),
	}
}

type GroupItems[I any] struct {
	Group string
	Items []I
}

func (o *GroupItems[I]) Count() int {
	return len(o.Items)
}

func (o *GroupItems[I]) ContainsItemBy(predicate func(item I) bool) (ret bool) {
	ret = lo.ContainsBy(o.Items, predicate)
	return
}

type GroupsItemsSelector[I any] struct {
	SelectionLabel string
	GetItemKey     func(I) string

	GroupsItems []*GroupItems[I]
}

func (o *GroupsItemsSelector[I]) AddGroupItems(group string, items ...I) {
	o.GroupsItems = append(o.GroupsItems, &GroupItems[I]{group, items})
}

// getSortedGroupsItems returns a new slice of GroupItems with both groups and their items
// sorted alphabetically in a case-insensitive manner. The original GroupsItems are not modified.
func (o *GroupsItemsSelector[I]) getSortedGroupsItems() []*GroupItems[I] {
	// Copy and sort groups (case‑insensitive)
	sortedGroupsItems := make([]*GroupItems[I], len(o.GroupsItems))
	copy(sortedGroupsItems, o.GroupsItems)
	sort.SliceStable(sortedGroupsItems, func(i, j int) bool {
		return strings.ToLower(sortedGroupsItems[i].Group) < strings.ToLower(sortedGroupsItems[j].Group)
	})

	// For each group, sort its items
	for i, groupItems := range sortedGroupsItems {
		sortedItems := make([]I, len(groupItems.Items))
		copy(sortedItems, groupItems.Items)
		sort.SliceStable(sortedItems, func(i, j int) bool {
			return strings.ToLower(o.GetItemKey(sortedItems[i])) < strings.ToLower(o.GetItemKey(sortedItems[j]))
		})

		// Create a new GroupItems with the sorted items
		sortedGroupsItems[i] = &GroupItems[I]{
			Group: groupItems.Group,
			Items: sortedItems,
		}
	}

	return sortedGroupsItems
}

func (o *GroupsItemsSelector[I]) GetGroupAndItemByItemNumber(number int) (group string, item I, err error) {
	var currentItemNumber int
	found := false

	sortedGroupsItems := o.getSortedGroupsItems()

	for _, groupItems := range sortedGroupsItems {
		if currentItemNumber+len(groupItems.Items) < number {
			currentItemNumber += len(groupItems.Items)
			continue
		}

		for _, groupItem := range groupItems.Items {
			currentItemNumber++
			if currentItemNumber == number {
				group = groupItems.Group
				item = groupItem
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	if !found {
		err = fmt.Errorf("number %d is out of range", number)
	}
	return
}

func (o *GroupsItemsSelector[I]) Print(shellCompleteList bool) {
	// Only print the section header if not in plain output mode
	if !shellCompleteList {
		fmt.Printf("\n%v:\n", o.SelectionLabel)
	}

	var currentItemIndex int
	sortedGroupsItems := o.getSortedGroupsItems()

	for _, groupItems := range sortedGroupsItems {
		if !shellCompleteList {
			fmt.Println()
			fmt.Printf("%s\n\n", groupItems.Group)
		}

		for _, item := range groupItems.Items {
			currentItemIndex++
			if shellCompleteList {
				// plain mode: "index key"
				fmt.Printf("%s\n", o.GetItemKey(item))
			} else {
				// formatted mode: "[index]    key"
				fmt.Printf("\t[%d]\t%s\n", currentItemIndex, o.GetItemKey(item))
			}
		}
	}
}

func (o *GroupsItemsSelector[I]) HasGroup(group string) (ret bool) {
	for _, groupItems := range o.GroupsItems {
		if ret = groupItems.Group == group; ret {
			break
		}
	}
	return
}

func (o *GroupsItemsSelector[I]) FindGroupsByItemFirst(item I) (ret string) {
	itemKey := o.GetItemKey(item)

	for _, groupItems := range o.GroupsItems {
		if groupItems.ContainsItemBy(func(groupItem I) bool {
			groupItemKey := o.GetItemKey(groupItem)
			return groupItemKey == itemKey
		}) {
			ret = groupItems.Group
			break
		}
	}
	return
}

func (o *GroupsItemsSelector[I]) FindGroupsByItem(item I) (groups []string) {
	itemKey := o.GetItemKey(item)

	for _, groupItems := range o.GroupsItems {
		if groupItems.ContainsItemBy(func(groupItem I) bool {
			groupItemKey := o.GetItemKey(groupItem)
			return groupItemKey == itemKey
		}) {
			groups = append(groups, groupItems.Group)
		}
	}
	return
}

func ReturnItem(item string) string {
	return item
}

func NewGroupsItemsSelectorString(selectionLabel string) *GroupsItemsSelectorString {
	return &GroupsItemsSelectorString{GroupsItemsSelector: NewGroupsItemsSelector(selectionLabel, ReturnItem)}
}

type GroupsItemsSelectorString struct {
	*GroupsItemsSelector[string]
}
