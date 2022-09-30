package view_models

type SectionsGetter interface {
	GetSections() []string
}

type HrefFormatterGetter interface {
	GetHrefFormatter() Formatter
}

type TitleFormatterGetter interface {
	GetTitleFormatter() Formatter
}

type ClassFormatterGetter interface {
	GetClassFormatter() Formatter
}

type ItemConfigurationProvider interface {
	PropertiesGetter
	ImagePropertyGetter
	ImagePathGetter
	SectionsGetter
	HrefFormatterGetter
	TitleFormatterGetter
	ClassFormatterGetter
}

type ItemConfigurationProviderGetter interface {
	GetItemConfigurationProvider() ItemConfigurationProvider
}
