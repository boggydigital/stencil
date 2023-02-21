package view_models

type ListDehydratedImagePropertyGetter interface {
	GetListDehydratedImageProperty() string
}

type ItemDehydratedImagePropertyGetter interface {
	GetItemDehydratedImageProperty() string
}

type DehydratedImagesConfigurator interface {
	ListDehydratedImagePropertyGetter
	ItemDehydratedImagePropertyGetter
}

type DehydratedImagesConfiguratorGetter interface {
	GetDehydratedImagesConfigurator() DehydratedImagesConfigurator
}
