package view_models

import (
	"github.com/boggydigital/kvas"
)

type Formatter func(id, property, link string, rxa kvas.ReduxAssets) string

type Item struct {
	*Page
	Id                string
	Title             string
	Labels            []string
	ImagePath         string
	ImageProperty     string
	TitleProperty     string
	SkippedProperties []string
	// Text properties
	Properties      map[string]map[string]string
	PropertyOrder   []string
	PropertyTitles  map[string]string
	PropertyClasses map[string]string
	// Icons
	Icons []string
	// Sections
	Sections      []string
	HasSections   []string
	SectionTitles map[string]string
}

func NewItem(
	acp AppConfigurationProvider,
	id string,
	hasSections []string,
	rxa kvas.ReduxAssets) (*Item, error) {

	icp := acp.GetItemConfigurationProvider()
	ccp := acp.GetCommonConfigurationProvider()

	if hasSections == nil {
		hasSections = icp.GetSections()
	}

	if err := rxa.IsSupported(icp.GetProperties()...); err != nil {
		return nil, err
	}

	title, _ := rxa.GetFirstVal(ccp.GetTitleProperty(), id)

	propertyOrder := append(icp.GetProperties(), icp.GetComputedProperties()...)

	ivm := &Item{
		Page:              acp.GetPage(),
		Id:                id,
		ImagePath:         icp.GetImagePath(),
		ImageProperty:     icp.GetImageProperty(),
		TitleProperty:     ccp.GetTitleProperty(),
		SkippedProperties: icp.GetSkippedProperties(),
		Title:             title,
		Labels:            ccp.GetLabels(),
		Properties:        make(map[string]map[string]string),
		PropertyOrder:     propertyOrder,
		PropertyTitles:    ccp.GetPropertyTitles(),
		PropertyClasses:   make(map[string]string),
		Icons:             ccp.GetIcons(),
		Sections:          icp.GetSections(),
		HasSections:       hasSections,
		SectionTitles:     ccp.GetSectionTitles(),
	}

	for _, p := range propertyOrder {
		ivm.Properties[p] = getPropertyLinks(
			id,
			p,
			icp.GetTitleFormatter(),
			icp.GetHrefFormatter(),
			rxa)
		if gcf := icp.GetClassFormatter(); gcf != nil {
			value, _ := rxa.GetFirstVal(p, id)
			ivm.PropertyClasses[p] = gcf(id, p, value, rxa)
		}
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
		linkTitle := fmtTitle(id, property, value, rxa)
		propertyLinks[linkTitle] = fmtHref(id, property, value, rxa)
	}

	return propertyLinks
}
