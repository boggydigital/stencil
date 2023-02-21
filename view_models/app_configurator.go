package view_models

type PageGetter interface {
	GetPage() *Page
}

type AppConfigurator interface {
	PageGetter
	CommonConfiguratorGetter
	ListConfiguratorGetter
	ItemConfiguratorGetter
	FormatterConfiguratorGetter
	SearchConfiguratorGetter
	DehydratedImagesConfiguratorGetter
}
