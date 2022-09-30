package view_models

import (
	"github.com/boggydigital/kvas"
)

type Formatter func(id, property, link string, rxa kvas.ReduxAssets) string

type Item struct {
	*Page
	Id            string
	Title         string
	Labels        []string
	ImagePath     string
	ImageProperty string
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

	icp := acp.GetItemConfigurationProvider()
	ccp := acp.GetCommonConfigurationProvider()

	if err := rxa.IsSupported(icp.GetProperties()...); err != nil {
		return nil, err
	}

	title, _ := rxa.GetFirstVal(ccp.GetTitleProperty(), id)

	ivm := &Item{
		Page:           acp.GetPage(),
		Id:             id,
		ImagePath:      icp.GetImagePath(),
		ImageProperty:  icp.GetImageProperty(),
		Title:          title,
		Labels:         ccp.GetLabels(),
		Properties:     make(map[string]map[string]string),
		PropertyOrder:  icp.GetProperties(),
		PropertyTitles: ccp.GetPropertyTitles(),
		Sections:       icp.GetSections(),
		SectionTitles:  ccp.GetSectionTitles(),
	}

	for _, p := range icp.GetProperties() {
		ivm.Properties[p] = getPropertyLinks(
			id,
			p,
			icp.GetTitleFormatter(),
			icp.GetHrefFormatter(),
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
		linkTitle := fmtTitle(id, property, value, rxa)
		propertyLinks[linkTitle] = fmtHref(id, property, value, rxa)
	}

	return propertyLinks
}
