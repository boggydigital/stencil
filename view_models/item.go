package view_models

import (
	"github.com/boggydigital/kvas"
)

type Formatter func(id, property, link string, rxa kvas.ReduxAssets) string

type Item struct {
	*Page
	Id               string
	Title            string
	Labels           []string
	HiddenLabels     []string
	LabelValues      map[string]string
	ImagePath        string
	ImageProperty    string
	TitleProperty    string
	HiddenProperties []string
	// Text properties
	Properties      map[string]map[string]string
	PropertyOrder   []string
	PropertyTitles  map[string]string
	PropertyClasses map[string]string
	PropertyActions map[string]map[string]string
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
	fcp := acp.GetFormatterConfigurationProvider()
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
		Page:             acp.GetPage(),
		Id:               id,
		ImagePath:        icp.GetImagePath(),
		ImageProperty:    icp.GetImageProperty(),
		TitleProperty:    ccp.GetTitleProperty(),
		Title:            title,
		Labels:           ccp.GetLabels(),
		HiddenLabels:     ccp.GetHiddenLabels(),
		LabelValues:      make(map[string]string),
		Properties:       make(map[string]map[string]string),
		PropertyOrder:    propertyOrder,
		HiddenProperties: icp.GetHiddenProperties(),
		PropertyTitles:   ccp.GetPropertyTitles(),
		PropertyClasses:  make(map[string]string),
		PropertyActions:  make(map[string]map[string]string),
		Icons:            ccp.GetIcons(),
		Sections:         icp.GetSections(),
		HasSections:      hasSections,
		SectionTitles:    ccp.GetSectionTitles(),
	}

	for _, p := range propertyOrder {
		value, _ := rxa.GetFirstVal(p, id)
		ivm.Properties[p] = getPropertyLinks(
			id,
			p,
			fcp.GetTitleFormatter(),
			fcp.GetHrefFormatter(),
			rxa)
		if gcf := fcp.GetClassFormatter(); gcf != nil {
			ivm.PropertyClasses[p] = gcf(id, p, value, rxa)
		}
		if pa := getPropertyActions(
			id,
			p,
			value,
			fcp.GetActionFormatter(),
			fcp.GetActionHrefFormatter(),
			rxa); len(pa) > 0 {
			ivm.PropertyActions[p] = pa
		}
	}

	glf := fcp.GetLabelFormatter()

	for _, l := range ccp.GetLabels() {
		if value, ok := rxa.GetFirstVal(l, id); ok {
			ivm.LabelValues[l] = glf(id, l, value, rxa)
		}
	}

	return ivm, nil
}

func getPropertyLinks(
	id, property string,
	fmtTitle, fmtHref Formatter,
	rxa kvas.ReduxAssets) map[string]string {

	propertyLinks := make(map[string]string)

	values, _ := rxa.GetAllUnchangedValues(property, id)

	for _, value := range values {
		linkTitle := fmtTitle(id, property, value, rxa)
		if linkTitle == "" {
			continue
		}
		propertyLinks[linkTitle] = fmtHref(id, property, value, rxa)
	}

	return propertyLinks
}

func getPropertyActions(
	id, property, value string,
	fmtAction, fmtActionHref Formatter,
	rxa kvas.ReduxAssets) map[string]string {

	actions := make(map[string]string)

	if a := fmtAction(id, property, value, rxa); a != "" {
		actions[a] = fmtActionHref(id, property, a, rxa)
	}

	return actions
}
