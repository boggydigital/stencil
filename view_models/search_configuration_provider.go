package view_models

type ScopesGetter interface {
	GetScopes() []string
}

type ScopeQueriesGetter interface {
	GetScopeQueries() map[string]string
}

type SearchConfigurationProvider interface {
	PropertiesGetter
	ScopesGetter
	ScopeQueriesGetter
}

type SearchConfigurationProviderGetter interface {
	GetSearchConfigurationProvider() SearchConfigurationProvider
}
