package view_models

type ComputedPropertiesGetter interface {
	GetComputedProperties() []string
}

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
	ComputedPropertiesGetter
	ImagePropertyGetter
	ImagePathGetter
	SectionsGetter
	SkippedPropertiesGetter
	HrefFormatterGetter
	TitleFormatterGetter
	ClassFormatterGetter
}

type ItemConfigurationProviderGetter interface {
	GetItemConfigurationProvider() ItemConfigurationProvider
}
