package stencil

import "github.com/boggydigital/stencil/view_models"

type ItemConfiguration struct {
	properties      []string
	imageProperty   string
	classProperties []string
	imagePath       string
	sections        []string
	hrefFormatter   view_models.Formatter
	titleFormatter  view_models.Formatter
	classFormatter  view_models.Formatter
}

func (ic *ItemConfiguration) GetProperties() []string {
	return ic.properties
}

func (ic *ItemConfiguration) GetImageProperty() string {
	return ic.imageProperty
}

func (ic *ItemConfiguration) GetClassProperties() []string {
	return ic.classProperties
}

func (ic *ItemConfiguration) GetImagePath() string {
	return ic.imagePath
}

func (ic *ItemConfiguration) GetSections() []string {
	return ic.sections
}

func (ic *ItemConfiguration) GetHrefFormatter() view_models.Formatter {
	return ic.hrefFormatter
}

func (ic *ItemConfiguration) GetTitleFormatter() view_models.Formatter {
	return ic.titleFormatter
}

func (ic *ItemConfiguration) GetClassFormatter() view_models.Formatter {
	return ic.classFormatter
}
