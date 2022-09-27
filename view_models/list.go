package view_models

import (
	"github.com/boggydigital/kvas"
)

type ListItem struct {
	Id              string
	Title           string
	Properties      []string
	LabelValues     map[string]string
	PropertyValues  map[string][]string
	PropertyClasses map[string]string
	PropertyTitles  map[string]string
}

type List struct {
	*Page
	Labels          []string
	Icons           []string
	ClassProperties []string
	ItemPath        string
	Items           []*ListItem
	TitleProperty   string
	CoverProperty   string
	CoverPath       string
}

func NewList(
	acp AppConfigurationProvider,
	ids []string,
	rxa kvas.ReduxAssets) (*List, error) {

	if err := rxa.IsSupported(acp.GetListProperties()...); err != nil {
		return nil, err
	}

	lvm := &List{
		Page:            acp.GetPage(),
		Labels:          acp.GetLabels(),
		Icons:           acp.GetIcons(),
		ClassProperties: acp.GetListClassProperties(),
		ItemPath:        acp.GetListItemPath(),
		Items:           make([]*ListItem, 0, len(ids)),
		TitleProperty:   acp.GetTitleProperty(),
		CoverProperty:   acp.GetListCoverProperty(),
		CoverPath:       acp.GetCoverPath(),
	}

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(acp.GetTitleProperty(), id)

		li := &ListItem{
			Id:              id,
			Title:           title,
			Properties:      acp.GetListProperties(),
			PropertyValues:  make(map[string][]string, len(acp.GetListProperties())),
			PropertyTitles:  acp.GetPropertyTitles(),
			LabelValues:     make(map[string]string, len(acp.GetLabels())),
			PropertyClasses: make(map[string]string, len(acp.GetLabels())),
		}

		for _, p := range acp.GetListProperties() {
			values, _ := rxa.GetAllUnchangedValues(p, id)
			li.PropertyValues[p] = values
		}

		for _, l := range acp.GetLabels() {
			if value, ok := rxa.GetFirstVal(l, id); ok {
				li.LabelValues[l] = acp.GetItemTitleFormatter()(id, l, value, rxa)
				if class := acp.GetItemClassFormatter()(id, l, value, rxa); class != "" {
					li.PropertyClasses[l] = class
				}
			}
		}

		lvm.Items = append(lvm.Items, li)
	}

	return lvm, nil
}
