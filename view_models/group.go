package view_models

import (
	"github.com/boggydigital/kevlar"
	"net/url"
)

type Group struct {
	*Page
	GroupOrder      []string
	GroupLists      map[string]*List
	GroupTitles     map[string]string
	GroupShowAll    map[string]bool
	AnyGroupShowAll bool
	Updated         string
	ShowAllUrl      string
}

func NewGroup(
	acp AppConfigurator,
	groupOrder []string,
	groupItems map[string][]string,
	groupTitles map[string]string,
	groupTotals map[string]int,
	updated string,
	u *url.URL,
	rdx kevlar.ReadableRedux) (*Group, error) {

	gvm := &Group{
		Page:        acp.GetPage(),
		GroupOrder:  groupOrder,
		GroupLists:  make(map[string]*List),
		GroupTitles: groupTitles,
		Updated:     updated,
	}

	showAll := make(map[string]bool)

	for group, items := range groupItems {
		lvm, err := NewList(acp, items, 0, len(items), groupTotals[group], nil, rdx)
		if err != nil {
			return gvm, err
		}
		gvm.GroupLists[group] = lvm
		if len(items) < groupTotals[group] {
			showAll[group] = true
		}
	}

	if u != nil && len(showAll) > 0 {
		q := u.Query()
		q.Set("show-all", "true")
		u.RawQuery = q.Encode()
		gvm.ShowAllUrl = u.String()
	}

	gvm.GroupShowAll = showAll
	gvm.AnyGroupShowAll = len(showAll) > 0

	return gvm, nil
}
