package view_models

type DehydratedImagePropertyGetter interface {
	GetDehydratedImageProperty() string
}

type DehydratedImagesConfigurator interface {
	DehydratedImagePropertyGetter
}

type DehydratedImagesConfiguratorGetter interface {
	GetDehydratedImagesConfigurator() DehydratedImagesConfigurator
}
