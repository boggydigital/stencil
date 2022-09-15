package view_models

import (
	"github.com/boggydigital/kvas"
)

type ListItem struct {
	Id         string
	Properties map[string][]string
}

type List struct {
	Items []*ListItem
}

func NewList(
	ids []string,
	properties []string,
	rxa kvas.ReduxAssets) (*List, error) {

	if err := rxa.IsSupported(properties...); err != nil {
		return nil, err
	}

	lvm := &List{
		Items: make([]*ListItem, 0, len(ids)),
	}

	for _, id := range ids {

		li := &ListItem{
			Id:         id,
			Properties: make(map[string][]string, len(properties)),
		}

		for _, p := range properties {
			li.Properties[p], _ = rxa.GetAllUnchangedValues(p, id)
		}

		lvm.Items = append(lvm.Items, &ListItem{
			Id: id,
		})
	}

	return lvm, nil
}
