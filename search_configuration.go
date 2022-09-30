package stencil

type SearchConfiguration struct {
	properties   []string
	scopes       []string
	scopeQueries map[string]string
}

func (sc *SearchConfiguration) GetProperties() []string {
	return sc.properties
}

func (sc *SearchConfiguration) GetScopes() []string {
	return sc.scopes
}

func (sc *SearchConfiguration) GetScopeQueries() map[string]string {
	return sc.scopeQueries
}
