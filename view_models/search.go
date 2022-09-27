package view_models

import (
	"github.com/boggydigital/kvas"
	"net/url"
	"strings"
)

type Search struct {
	*Page
	Scopes         []string
	ScopeQueries   map[string]string
	CurrentScope   string
	Properties     []string
	PropertyTitles map[string]string
	Query          map[string][]string
	Digests        map[string][]string
	DigestsTitles  map[string]string
	Found          *List
}

func NewSearch(
	acp AppConfigurationProvider,
	query map[string][]string,
	ids []string,
	digests map[string][]string,
	rxa kvas.ReduxAssets) (*Search, error) {

	lvm, err := NewList(acp, ids, rxa)

	svm := &Search{
		Scopes:         acp.GetSearchScopes(),
		ScopeQueries:   acp.GetSearchScopeQueries(),
		CurrentScope:   currentScope(query, acp.GetSearchScopeQueries()),
		Query:          query,
		Page:           acp.GetPage(),
		Properties:     acp.GetSearchProperties(),
		PropertyTitles: acp.GetPropertyTitles(),
		Digests:        digests,
		DigestsTitles:  acp.GetDigestTitles(),
		Found:          lvm,
	}

	return svm, err
}

func currentScope(query map[string][]string, scopeUrls map[string]string) string {

	q := url.Values{}

	for k, vs := range query {
		q.Set(k, strings.Join(vs, ","))
	}

	if eq, err := url.QueryUnescape(q.Encode()); err == nil {
		for scope, scopeUrl := range scopeUrls {
			if scopeUrl == eq {
				return scope
			}
		}
	}

	return ""
}
