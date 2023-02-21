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

type HiddenPropertiesGetter interface {
	GetHiddenProperties() []string
}

type ListConfigurator interface {
	PropertiesGetter
	HiddenPropertiesGetter
	ImagePropertyGetter
	ImagePathGetter
	ItemPathGetter
}

type ListConfiguratorGetter interface {
	GetListConfigurator() ListConfigurator
}
