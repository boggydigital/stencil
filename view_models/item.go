package view_models

import (
	"github.com/boggydigital/kvas"
)

type Formatter func(property, link string) string

type Item struct {
	*Page
	Id        string
	Title     string
	CoverHref string
	// Text properties
	Properties      map[string]map[string]string
	PropertyOrder   []string
	PropertyTitles  map[string]string
	PropertyClasses map[string]string
	// Sections
	Sections      []string
	SectionTitles map[string]string
}

func NewItem(
	page *Page,
	id string,
	coverHref string,
	properties []string,
	titleProperty string,
	propertyTitles map[string]string,
	sections []string,
	sectionTitles map[string]string,
	fmtTitle, fmtHref Formatter,
	rxa kvas.ReduxAssets) (*Item, error) {

	if err := rxa.IsSupported(properties...); err != nil {
		return nil, err
	}

	title, _ := rxa.GetFirstVal(titleProperty, id)

	ivm := &Item{
		Page:           page,
		Id:             id,
		CoverHref:      coverHref,
		Title:          title,
		Properties:     make(map[string]map[string]string),
		PropertyOrder:  properties,
		PropertyTitles: propertyTitles,
		Sections:       sections,
		SectionTitles:  sectionTitles,
	}

	for _, p := range properties {
		ivm.Properties[p] = getPropertyLinks(id, p, fmtTitle, fmtHref, rxa)
	}

	return ivm, nil
}

func getPropertyLinks(
	id string,
	property string,
	fmtTitle, fmtHref Formatter,
	rxa kvas.ReduxAssets) map[string]string {

	propertyLinks := make(map[string]string)

	values, _ := rxa.GetAllUnchangedValues(property, id)

	for _, value := range values {
		linkTitle := fmtTitle(property, value)
		propertyLinks[linkTitle] = fmtHref(property, value)
	}

	return propertyLinks
}
