package view_models

type PageGetter interface {
	GetPage() *Page
}

type AppConfigurationProvider interface {
	PageGetter
	CommonConfigurationProviderGetter
	ListConfigurationProviderGetter
	ItemConfigurationProviderGetter
	SearchConfigurationProviderGetter
}
