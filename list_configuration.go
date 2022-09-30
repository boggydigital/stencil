package stencil

type ListConfiguration struct {
	properties      []string
	itemPath        string
	imageProperty   string
	imagePath       string
	classProperties []string
}

func (lc *ListConfiguration) GetProperties() []string {
	return lc.properties
}

func (lc *ListConfiguration) GetItemPath() string {
	return lc.itemPath
}

func (lc *ListConfiguration) GetImageProperty() string {
	return lc.imageProperty
}

func (lc *ListConfiguration) GetImagePath() string {
	return lc.imagePath
}

func (lc *ListConfiguration) GetClassProperties() []string {
	return lc.classProperties
}
