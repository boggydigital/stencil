package view_models

type ScopesGetter interface {
	GetScopes() []string
}

type DigestPropertiesGetter interface {
	GetDigestProperties() []string
}

type HighlightPropertiesGetter interface {
	GetHighlightProperties() []string
}

type ScopeQueriesGetter interface {
	GetScopeQueries() map[string]string
}

type SearchConfigurationProvider interface {
	PropertiesGetter
	HighlightPropertiesGetter
	DigestPropertiesGetter
	ScopesGetter
	ScopeQueriesGetter
}

type SearchConfigurationProviderGetter interface {
	GetSearchConfigurationProvider() SearchConfigurationProvider
}
