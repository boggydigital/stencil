package view_models

type ScopesGetter interface {
	GetScopes() []string
}

type DigestPropertiesGetter interface {
	GetDigestProperties() []string
}

type ScopeQueriesGetter interface {
	GetScopeQueries() map[string]string
}

type SearchConfigurationProvider interface {
	PropertiesGetter
	DigestPropertiesGetter
	ScopesGetter
	ScopeQueriesGetter
}

type SearchConfigurationProviderGetter interface {
	GetSearchConfigurationProvider() SearchConfigurationProvider
}
