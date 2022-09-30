package stencil

type CommonConfiguration struct {
	labels         []string
	icons          []string
	titleProperty  string
	propertyTitles map[string]string
	sectionTitles  map[string]string
	digestTitles   map[string]string
}

func (cc *CommonConfiguration) GetLabels() []string {
	return cc.labels
}

func (cc *CommonConfiguration) GetIcons() []string {
	return cc.icons
}

func (cc *CommonConfiguration) GetTitleProperty() string {
	return cc.titleProperty
}

func (cc *CommonConfiguration) GetPropertyTitles() map[string]string {
	return cc.propertyTitles
}

func (cc *CommonConfiguration) GetSectionTitles() map[string]string {
	return cc.sectionTitles
}

func (cc *CommonConfiguration) GetDigestTitles() map[string]string {
	return cc.digestTitles
}
