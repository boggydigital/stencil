package view_models

import (
	"github.com/boggydigital/kvas"
)

const eagerLoadingCovers = 10

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
	Labels             []string
	Icons              []string
	ClassProperties    []string
	ItemPath           string
	Items              []*ListItem
	TitleProperty      string
	ImageProperty      string
	ImagePath          string
	EagerLoadingCovers int
}

func NewList(
	acp AppConfigurationProvider,
	ids []string,
	rxa kvas.ReduxAssets) (*List, error) {

	lcp := acp.GetListConfigurationProvider()
	ccp := acp.GetCommonConfigurationProvider()

	if err := rxa.IsSupported(lcp.GetProperties()...); err != nil {
		return nil, err
	}

	lvm := &List{
		Page:               acp.GetPage(),
		Labels:             ccp.GetLabels(),
		Icons:              ccp.GetIcons(),
		ClassProperties:    lcp.GetClassProperties(),
		ItemPath:           lcp.GetItemPath(),
		Items:              make([]*ListItem, 0, len(ids)),
		TitleProperty:      ccp.GetTitleProperty(),
		ImageProperty:      lcp.GetImageProperty(),
		ImagePath:          lcp.GetImagePath(),
		EagerLoadingCovers: eagerLoadingCovers,
	}

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(ccp.GetTitleProperty(), id)

		li := &ListItem{
			Id:              id,
			Title:           title,
			Properties:      lcp.GetProperties(),
			PropertyValues:  make(map[string][]string, len(lcp.GetProperties())),
			PropertyTitles:  ccp.GetPropertyTitles(),
			LabelValues:     make(map[string]string, len(ccp.GetLabels())),
			PropertyClasses: make(map[string]string, len(ccp.GetLabels())),
		}

		for _, p := range lcp.GetProperties() {
			values, _ := rxa.GetAllUnchangedValues(p, id)
			li.PropertyValues[p] = values
		}

		icp := acp.GetItemConfigurationProvider()

		for _, l := range ccp.GetLabels() {
			if value, ok := rxa.GetFirstVal(l, id); ok {
				li.LabelValues[l] = icp.GetTitleFormatter()(id, l, value, rxa)
				if icp.GetClassFormatter() != nil {
					if class := icp.GetClassFormatter()(id, l, value, rxa); class != "" {
						li.PropertyClasses[l] = class
					}
				}
			}
		}

		lvm.Items = append(lvm.Items, li)
	}

	return lvm, nil
}
