package view_models

import (
	"github.com/boggydigital/kvas"
)

type Formatter func(id, property, link string) string

type Item struct {
	*Page
	Id        string
	Title     string
	Labels    []string
	CoverPath string
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
	acp AppConfigurationProvider,
	id string,
	rxa kvas.ReduxAssets) (*Item, error) {

	if err := rxa.IsSupported(acp.GetItemProperties()...); err != nil {
		return nil, err
	}

	title, _ := rxa.GetFirstVal(acp.GetTitleProperty(), id)

	ivm := &Item{
		Page:           acp.GetPage(),
		Id:             id,
		CoverPath:      acp.GetCoverPath(),
		Title:          title,
		Labels:         acp.GetLabels(),
		Properties:     make(map[string]map[string]string),
		PropertyOrder:  acp.GetItemProperties(),
		PropertyTitles: acp.GetPropertyTitles(),
		Sections:       acp.GetItemSections(),
		SectionTitles:  acp.GetSectionTitles(),
	}

	for _, p := range acp.GetItemProperties() {
		ivm.Properties[p] = getPropertyLinks(
			id,
			p,
			acp.GetItemTitleFormatter(),
			acp.GetItemHrefFormatter(),
			rxa)
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
		linkTitle := fmtTitle(id, property, value)
		propertyLinks[linkTitle] = fmtHref(id, property, value)
	}

	return propertyLinks
}
