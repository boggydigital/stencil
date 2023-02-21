package view_models

import (
	"github.com/boggydigital/kvas"
)

type Formatter func(id, property, link string, rxa kvas.ReduxAssets) string

type Item struct {
	*Page
	Id                      string
	Title                   string
	Labels                  []string
	HiddenLabels            []string
	LabelValues             map[string]string
	ImagePath               string
	DehydratedImageProperty string
	ImageProperty           string
	TitleProperty           string
	HiddenProperties        []string
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
	acp AppConfigurator,
	id string,
	hasSections []string,
	rxa kvas.ReduxAssets) (*Item, error) {

	ic := acp.GetItemConfigurator()
	fc := acp.GetFormatterConfigurator()
	cc := acp.GetCommonConfigurator()
	dic := acp.GetDehydratedImagesConfigurator()

	if hasSections == nil {
		hasSections = ic.GetSections()
	}

	if err := rxa.IsSupported(ic.GetProperties()...); err != nil {
		return nil, err
	}

	title, _ := rxa.GetFirstVal(cc.GetTitleProperty(), id)

	propertyOrder := append(ic.GetProperties(), ic.GetComputedProperties()...)

	ivm := &Item{
		Page:                    acp.GetPage(),
		Id:                      id,
		ImagePath:               ic.GetImagePath(),
		DehydratedImageProperty: dic.GetItemDehydratedImageProperty(),
		ImageProperty:           ic.GetImageProperty(),
		TitleProperty:           cc.GetTitleProperty(),
		Title:                   title,
		Labels:                  cc.GetLabels(),
		HiddenLabels:            cc.GetHiddenLabels(),
		LabelValues:             make(map[string]string),
		Properties:              make(map[string]map[string]string),
		PropertyOrder:           propertyOrder,
		HiddenProperties:        ic.GetHiddenProperties(),
		PropertyTitles:          cc.GetPropertyTitles(),
		PropertyClasses:         make(map[string]string),
		PropertyActions:         make(map[string]map[string]string),
		Icons:                   cc.GetIcons(),
		Sections:                ic.GetSections(),
		HasSections:             hasSections,
		SectionTitles:           cc.GetSectionTitles(),
	}

	for _, p := range propertyOrder {
		value, _ := rxa.GetFirstVal(p, id)
		ivm.Properties[p] = getPropertyLinks(
			id,
			p,
			fc.GetTitleFormatter(),
			fc.GetHrefFormatter(),
			rxa)
		if gcf := fc.GetClassFormatter(); gcf != nil {
			ivm.PropertyClasses[p] = gcf(id, p, value, rxa)
		}
		if pa := getPropertyActions(
			id,
			p,
			value,
			fc.GetActionFormatter(),
			fc.GetActionHrefFormatter(),
			rxa); len(pa) > 0 {
			ivm.PropertyActions[p] = pa
		}
	}

	glf := fc.GetLabelFormatter()

	for _, l := range cc.GetLabels() {
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

	if fmtAction == nil {
		return nil
	}

	actions := make(map[string]string)

	if a := fmtAction(id, property, value, rxa); a != "" {
		actions[a] = fmtActionHref(id, property, a, rxa)
	}

	return actions
}
