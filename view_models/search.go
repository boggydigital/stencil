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
	from, to, total int,
	u *url.URL,
	digests map[string][]string,
	rxa kvas.ReduxAssets) (*Search, error) {

	lvm, err := NewList(acp, ids, from, to, total, u, rxa)

	scp := acp.GetSearchConfigurationProvider()
	ccp := acp.GetCommonConfigurationProvider()

	svm := &Search{
		Scopes:         scp.GetScopes(),
		ScopeQueries:   scp.GetScopeQueries(),
		CurrentScope:   currentScope(query, scp.GetScopeQueries()),
		Query:          query,
		Page:           acp.GetPage(),
		Properties:     scp.GetProperties(),
		PropertyTitles: ccp.GetPropertyTitles(),
		Digests:        digests,
		DigestsTitles:  ccp.GetDigestTitles(),
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
