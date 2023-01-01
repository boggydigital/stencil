package view_models

import (
	"github.com/boggydigital/kvas"
	"net/url"
	"strconv"
)

const eagerLoadingImages = 10

type ListItem struct {
	Id              string
	Title           string
	LabelValues     map[string]string
	PropertyValues  map[string][]string
	PropertyClasses map[string]string
	PropertyTitles  map[string]string
}

type List struct {
	*Page
	Properties         []string
	Labels             []string
	HiddenLabels       []string
	Icons              []string
	HiddenProperties   []string
	ItemPath           string
	Items              []*ListItem
	From, To, Total    int
	NextUrl            string
	TitleProperty      string
	ImageProperty      string
	ImagePath          string
	EagerLoadingImages int
}

func NewList(
	acp AppConfigurationProvider,
	ids []string,
	from, to, total int,
	u *url.URL,
	rxa kvas.ReduxAssets) (*List, error) {

	lcp := acp.GetListConfigurationProvider()
	ccp := acp.GetCommonConfigurationProvider()

	if err := rxa.IsSupported(lcp.GetProperties()...); err != nil {
		return nil, err
	}

	lvm := &List{
		Page:               acp.GetPage(),
		Properties:         lcp.GetProperties(),
		Labels:             ccp.GetLabels(),
		HiddenLabels:       ccp.GetHiddenLabels(),
		Icons:              ccp.GetIcons(),
		HiddenProperties:   lcp.GetHiddenProperties(),
		ItemPath:           lcp.GetItemPath(),
		Items:              make([]*ListItem, 0, len(ids)),
		From:               from + 1,
		To:                 to,
		Total:              total,
		TitleProperty:      ccp.GetTitleProperty(),
		ImageProperty:      lcp.GetImageProperty(),
		ImagePath:          lcp.GetImagePath(),
		EagerLoadingImages: eagerLoadingImages,
	}

	if u != nil &&
		to < total {
		q := u.Query()
		q.Set("from", strconv.Itoa(to))
		u.RawQuery = q.Encode()
		lvm.NextUrl = u.String()
	}

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(ccp.GetTitleProperty(), id)

		li := &ListItem{
			Id:              id,
			Title:           title,
			PropertyValues:  make(map[string][]string, len(lcp.GetProperties())),
			PropertyTitles:  ccp.GetPropertyTitles(),
			LabelValues:     make(map[string]string, len(ccp.GetLabels())),
			PropertyClasses: make(map[string]string, len(ccp.GetLabels())),
		}

		for _, p := range lcp.GetProperties() {
			values, _ := rxa.GetAllUnchangedValues(p, id)
			li.PropertyValues[p] = values
		}

		fcp := acp.GetFormatterConfigurationProvider()

		glf := fcp.GetLabelFormatter()
		gcf := fcp.GetClassFormatter()

		for _, l := range ccp.GetLabels() {
			if value, ok := rxa.GetFirstVal(l, id); ok {
				li.LabelValues[l] = glf(id, l, value, rxa)
				if gcf == nil {
					continue
				}
				if class := gcf(id, l, value, rxa); class != "" {
					li.PropertyClasses[l] = class
				}
			}
		}

		lvm.Items = append(lvm.Items, li)
	}

	return lvm, nil
}
