package view_models

import (
	"github.com/boggydigital/kvas"
	"strings"
)

type ListItem struct {
	Id             string
	Properties     []string
	PropertyValues map[string]string
}

type List struct {
	Nav          *Navigation
	AppTemplates []string
	ItemHref     string
	Items        []*ListItem
}

func NewList(
	nav *Navigation,
	itemHref string,
	ids []string,
	properties []string,
	rxa kvas.ReduxAssets) (*List, error) {

	if err := rxa.IsSupported(properties...); err != nil {
		return nil, err
	}

	lvm := &List{
		Nav:      nav,
		ItemHref: itemHref,
		Items:    make([]*ListItem, 0, len(ids)),
	}

	for _, id := range ids {

		li := &ListItem{
			Id:             id,
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
