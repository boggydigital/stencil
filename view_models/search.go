package view_models

import "github.com/boggydigital/kvas"

type Search struct {
	*Page
	Properties     []string
	PropertyTitles map[string]string
	Query          map[string]string
	Digests        map[string][]string
	DigestsTitles  map[string]string
	List           *List
}

func NewSearch(
	page *Page,
	itemHref string,
	ids []string,
	searchProperties []string,
	listProperties []string,
	propertyTitles map[string]string,
	digestTitles map[string]string,
	rxa kvas.ReduxAssets) (*Search, error) {

	lvm, err := NewList(page, itemHref, ids, listProperties, rxa)

	svm := &Search{
		Page:           page,
		Properties:     searchProperties,
		PropertyTitles: propertyTitles,
		Query:          make(map[string]string),
		DigestsTitles:  digestTitles,
		List:           lvm,
	}

	return svm, err
}
