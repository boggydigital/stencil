package view_models

import (
	"github.com/boggydigital/kvas"
	"net/url"
	"strings"
)

type Search struct {
	*Page
	Scopes           []string
	ScopeQueries     map[string]string
	CurrentScope     string
	Properties       []string
	PropertyTitles   map[string]string
	Query            map[string][]string
	DigestProperties []string
	Found            *List
}

func NewSearch(
	acp AppConfigurator,
	query map[string][]string,
	ids []string,
	from, to, total int,
	u *url.URL,
	//digests map[string][]string,
	rdx kvas.ReadableRedux) (*Search, error) {

	lvm, err := NewList(acp, ids, from, to, total, u, rdx)

	sc := acp.GetSearchConfigurator()
	cc := acp.GetCommonConfigurator()

	svm := &Search{
		Scopes:           sc.GetScopes(),
		ScopeQueries:     sc.GetScopeQueries(),
		CurrentScope:     currentScope(query, sc.GetScopeQueries()),
		Query:            query,
		Page:             acp.GetPage(),
		Properties:       sc.GetProperties(),
		PropertyTitles:   cc.GetPropertyTitles(),
		DigestProperties: sc.GetDigestProperties(),
		Found:            lvm,
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
