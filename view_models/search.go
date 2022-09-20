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
	page *Page,
	navPath string,
	scopes []string,
	scopeQueries map[string]string,
	query map[string][]string,
	ids []string,
	searchProperties []string,
	titleProperty string,
	labels []string,
	listProperties []string,
	propertyTitles map[string]string,
	digests map[string][]string,
	digestTitles map[string]string,
	fmtTitle Formatter,
	rxa kvas.ReduxAssets) (*Search, error) {

	lvm, err := NewList(
		page,
		navPath,
		ids,
		titleProperty,
		labels,
		listProperties,
		propertyTitles,
		fmtTitle,
		rxa)

	svm := &Search{
		Scopes:         scopes,
		ScopeQueries:   scopeQueries,
		CurrentScope:   currentScope(query, scopeQueries),
		Query:          query,
		Page:           page,
		Properties:     searchProperties,
		PropertyTitles: propertyTitles,
		Digests:        digests,
		DigestsTitles:  digestTitles,
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
