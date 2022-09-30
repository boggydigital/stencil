package view_models

type LabelsGetter interface {
	GetLabels() []string
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

type DigestTitlesGetter interface {
	GetDigestTitles() map[string]string
}

type CommonConfigurationProvider interface {
	LabelsGetter
	IconsGetter
	TitlePropertyGetter
	PropertyTitlesGetter
	SectionTitlesGetter
	DigestTitlesGetter
}

type CommonConfigurationProviderGetter interface {
	GetCommonConfigurationProvider() CommonConfigurationProvider
}
