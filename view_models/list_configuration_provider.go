package view_models

type PropertiesGetter interface {
	GetProperties() []string
}

type ItemPathGetter interface {
	GetItemPath() string
}

type ImagePropertyGetter interface {
	GetImageProperty() string
}

type ImagePathGetter interface {
	GetImagePath() string
}

type ClassPropertiesGetter interface {
	GetClassProperties() []string
}

type ListConfigurationProvider interface {
	PropertiesGetter
	ImagePropertyGetter
	ImagePathGetter
	ClassPropertiesGetter
	ItemPathGetter
}

type ListConfigurationProviderGetter interface {
	GetListConfigurationProvider() ListConfigurationProvider
}
