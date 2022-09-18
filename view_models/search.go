package view_models

import "github.com/boggydigital/kvas"

type Search struct {
	*Page
	Properties     []string
	PropertyTitles map[string]string
	Query          map[string][]string
	Digests        map[string][]string
	DigestsTitles  map[string]string
	List           *List
}

func NewSearch(
	page *Page,
	itemHref string,
	query map[string][]string,
	ids []string,
	searchProperties []string,
	titleProperty string,
	labels []string,
	listProperties []string,
	propertyTitles map[string]string,
	digests map[string][]string,
	digestTitles map[string]string,
	rxa kvas.ReduxAssets) (*Search, error) {

	lvm, err := NewList(page, itemHref, ids, titleProperty, labels, listProperties, rxa)

	svm := &Search{
		Query:          query,
		Page:           page,
		Properties:     searchProperties,
		PropertyTitles: propertyTitles,
		Digests:        digests,
		DigestsTitles:  digestTitles,
		List:           lvm,
	}

	return svm, err
}
