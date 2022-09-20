package view_models

import "github.com/boggydigital/kvas"

type Group struct {
	*Page
	GroupOrder  []string
	GroupLists  map[string]*List
	GroupTitles map[string]string
	Updated     string
}

func NewGroup(
	page *Page,
	itemPath string,
	groupOrder []string,
	groupItems map[string][]string,
	groupTitles map[string]string,
	titleProperty string,
	labels []string,
	listProperties []string,
	propertyTitles map[string]string,
	fmtTitle Formatter,
	updated string,
	rxa kvas.ReduxAssets) (*Group, error) {

	gvm := &Group{
		Page:        page,
		GroupOrder:  groupOrder,
		GroupLists:  make(map[string]*List),
		GroupTitles: groupTitles,
		Updated:     updated,
	}

	for group, items := range groupItems {
		lvm, err := NewList(
			page,
			itemPath,
			items,
			titleProperty,
			labels,
			listProperties,
			propertyTitles,
			fmtTitle,
			rxa)
		if err != nil {
			return gvm, err
		}
		gvm.GroupLists[group] = lvm
	}

	return gvm, nil
}
