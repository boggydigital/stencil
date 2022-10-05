package stencil

import "github.com/boggydigital/stencil/view_models"

type FormatterConfiguration struct {
	titleFormatter  view_models.Formatter
	hrefFormatter   view_models.Formatter
	classFormatter  view_models.Formatter
	actionFormatter view_models.Formatter
}

func (fc *FormatterConfiguration) GetTitleFormatter() view_models.Formatter {
	return fc.titleFormatter
}

func (fc *FormatterConfiguration) GetClassFormatter() view_models.Formatter {
	return fc.classFormatter
}

func (fc *FormatterConfiguration) GetHrefFormatter() view_models.Formatter {
	return fc.hrefFormatter
}

func (fc *FormatterConfiguration) GetActionFormatter() view_models.Formatter {
	return fc.actionFormatter
}
