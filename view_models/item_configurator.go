package view_models

type ComputedPropertiesGetter interface {
	GetComputedProperties() []string
}

type SectionsGetter interface {
	GetSections() []string
}

type ItemConfigurator interface {
	PropertiesGetter
	ComputedPropertiesGetter
	HiddenPropertiesGetter
	ImagePropertyGetter
	ImagePathGetter
	SectionsGetter
}

type ItemConfiguratorGetter interface {
	GetItemConfigurator() ItemConfigurator
}
