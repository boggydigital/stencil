package stencil

type ItemConfiguration struct {
	properties         []string
	computedProperties []string
	hiddenProperties   []string
	imageProperty      string
	imagePath          string
	sections           []string
}

func (ic *ItemConfiguration) GetProperties() []string {
	return ic.properties
}

func (ic *ItemConfiguration) GetComputedProperties() []string {
	return ic.computedProperties
}

func (ic *ItemConfiguration) GetImageProperty() string {
	return ic.imageProperty
}

func (ic *ItemConfiguration) GetHiddenProperties() []string {
	return ic.hiddenProperties
}

func (ic *ItemConfiguration) GetImagePath() string {
	return ic.imagePath
}

func (ic *ItemConfiguration) GetSections() []string {
	return ic.sections
}
