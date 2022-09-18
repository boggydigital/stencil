package view_models

import (
	"github.com/boggydigital/kvas"
	"strings"
)

type ListItem struct {
	Id             string
	Title          string
	Properties     []string
	PropertyValues map[string]string
}

type List struct {
	*Page
	AppTemplates []string
	ItemHref     string
	Items        []*ListItem
}

func NewList(
	page *Page,
	itemHref string,
	ids []string,
	titleProperty string,
	properties []string,
	rxa kvas.ReduxAssets) (*List, error) {

	if err := rxa.IsSupported(properties...); err != nil {
		return nil, err
	}

	lvm := &List{
		Page:     page,
		ItemHref: itemHref,
		Items:    make([]*ListItem, 0, len(ids)),
	}

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(titleProperty, id)

		li := &ListItem{
			Id:             id,
			Title:          title,
			Properties:     properties,
			PropertyValues: make(map[string]string, len(properties)),
		}

		for _, p := range properties {
			values, _ := rxa.GetAllUnchangedValues(p, id)
			li.PropertyValues[p] = strings.Join(values, ", ")
		}

		lvm.Items = append(lvm.Items, li)
	}

	return lvm, nil
}
