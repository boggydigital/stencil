package view_models

import "github.com/boggydigital/kvas"

type Search struct {
	Properties    []string
	Query         map[string]string
	Digests       map[string][]string
	DigestsTitles map[string]string
	List          *List
}

func NewSearch(
	ids []string,
	searchProperties []string,
	listProperties []string,
	rxa kvas.ReduxAssets) (*Search, error) {

	lvm, err := NewList(ids, listProperties, rxa)

	svm := &Search{
		Properties: searchProperties,
		Query:      make(map[string]string),
		//DigestsTitles:    digestTitles,
		List: lvm,
	}

	return svm, err
}
