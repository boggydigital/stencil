package stencil

type SearchConfiguration struct {
	properties       []string
	digestProperties []string
	scopes           []string
	scopeQueries     map[string]string
}

func (sc *SearchConfiguration) GetProperties() []string {
	return sc.properties
}

func (sc *SearchConfiguration) GetDigestProperties() []string {
	return sc.digestProperties
}

func (sc *SearchConfiguration) GetScopes() []string {
	return sc.scopes
}

func (sc *SearchConfiguration) GetScopeQueries() map[string]string {
	return sc.scopeQueries
}
