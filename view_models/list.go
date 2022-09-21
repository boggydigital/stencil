package view_models

import (
	"github.com/boggydigital/kvas"
	"strings"
)

type ListItem struct {
	Id             string
	Title          string
	Properties     []string
	LabelValues    map[string]string
	PropertyValues map[string]string
	PropertyTitles map[string]string
}

type List struct {
	*Page
	Labels        []string
	ItemPath      string
	Items         []*ListItem
	TitleProperty string
}

func NewList(
	page *Page,
	itemPath string,
	ids []string,
	titleProperty string,
	labels []string,
	properties []string,
	propertyTitles map[string]string,
	fmtTitle Formatter,
	rxa kvas.ReduxAssets) (*List, error) {

	if err := rxa.IsSupported(properties...); err != nil {
		return nil, err
	}

	lvm := &List{
		Page:          page,
		Labels:        labels,
		ItemPath:      itemPath,
		Items:         make([]*ListItem, 0, len(ids)),
		TitleProperty: titleProperty,
	}

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(titleProperty, id)

		li := &ListItem{
			Id:             id,
			Title:          title,
			Properties:     properties,
			PropertyValues: make(map[string]string, len(properties)),
			PropertyTitles: propertyTitles,
			LabelValues:    make(map[string]string, len(labels)),
		}

		for _, p := range properties {
			values, _ := rxa.GetAllUnchangedValues(p, id)
			li.PropertyValues[p] = strings.Join(values, ", ")
		}

		for _, l := range labels {
			if value, ok := rxa.GetFirstVal(l, id); ok {
				li.LabelValues[l] = fmtTitle(id, l, value)
			}
		}

		lvm.Items = append(lvm.Items, li)
	}

	return lvm, nil
}
