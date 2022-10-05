package view_models

type ComputedPropertiesGetter interface {
	GetComputedProperties() []string
}

type SectionsGetter interface {
	GetSections() []string
}

type ItemConfigurationProvider interface {
	PropertiesGetter
	ComputedPropertiesGetter
	HiddenPropertiesGetter
	ImagePropertyGetter
	ImagePathGetter
	SectionsGetter
}

type ItemConfigurationProviderGetter interface {
	GetItemConfigurationProvider() ItemConfigurationProvider
}
