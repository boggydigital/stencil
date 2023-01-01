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
	acp AppConfigurationProvider,
	groupOrder []string,
	groupItems map[string][]string,
	groupTitles map[string]string,
	updated string,
	rxa kvas.ReduxAssets) (*Group, error) {

	gvm := &Group{
		Page:        acp.GetPage(),
		GroupOrder:  groupOrder,
		GroupLists:  make(map[string]*List),
		GroupTitles: groupTitles,
		Updated:     updated,
	}

	for group, items := range groupItems {
		lvm, err := NewList(acp, items, 0, len(items), len(items), nil, rxa)
		if err != nil {
			return gvm, err
		}
		gvm.GroupLists[group] = lvm
	}

	return gvm, nil
}
