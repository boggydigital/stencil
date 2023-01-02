package view_models

type LabelsGetter interface {
	GetLabels() []string
}

type HiddenLabelsGetter interface {
	GetHiddenLabels() []string
}

type IconsGetter interface {
	GetIcons() []string
}

type TitlePropertyGetter interface {
	GetTitleProperty() string
}

type PropertyTitlesGetter interface {
	GetPropertyTitles() map[string]string
}

type SectionTitlesGetter interface {
	GetSectionTitles() map[string]string
}

type CommonConfigurationProvider interface {
	LabelsGetter
	HiddenLabelsGetter
	IconsGetter
	TitlePropertyGetter
	PropertyTitlesGetter
	SectionTitlesGetter
}

type CommonConfigurationProviderGetter interface {
	GetCommonConfigurationProvider() CommonConfigurationProvider
}
