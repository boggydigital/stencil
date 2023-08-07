package view_models

import (
	"github.com/boggydigital/kvas"
	"net/url"
	"strconv"
)

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
	Properties              []string
	Labels                  []string
	HiddenLabels            []string
	Icons                   []string
	HiddenProperties        []string
	ItemPath                string
	Items                   []*ListItem
	From, To, Total         int
	NextUrl                 string
	TitleProperty           string
	DehydratedImageProperty string
	ImageProperty           string
	ImagePath               string
}

func NewList(
	acp AppConfigurator,
	ids []string,
	from, to, total int,
	u *url.URL,
	rxa kvas.ReduxAssets) (*List, error) {

	lc := acp.GetListConfigurator()
	cc := acp.GetCommonConfigurator()
	dic := acp.GetDehydratedImagesConfigurator()

	if err := rxa.IsSupported(lc.GetProperties()...); err != nil {
		return nil, err
	}

	lvm := &List{
		Page:                    acp.GetPage(),
		Properties:              lc.GetProperties(),
		Labels:                  cc.GetLabels(),
		HiddenLabels:            cc.GetHiddenLabels(),
		Icons:                   cc.GetIcons(),
		HiddenProperties:        lc.GetHiddenProperties(),
		ItemPath:                lc.GetItemPath(),
		Items:                   make([]*ListItem, 0, len(ids)),
		From:                    from + 1,
		To:                      to,
		Total:                   total,
		TitleProperty:           cc.GetTitleProperty(),
		DehydratedImageProperty: dic.GetListDehydratedImageProperty(),
		ImageProperty:           lc.GetImageProperty(),
		ImagePath:               lc.GetImagePath(),
	}

	if u != nil &&
		to < total {
		q := u.Query()
		q.Set("from", strconv.Itoa(to))
		u.RawQuery = q.Encode()
		lvm.NextUrl = u.String()
	}

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(cc.GetTitleProperty(), id)

		li := &ListItem{
			Id:              id,
			Title:           title,
			PropertyValues:  make(map[string][]string, len(lc.GetProperties())),
			PropertyTitles:  cc.GetPropertyTitles(),
			LabelValues:     make(map[string]string, len(cc.GetLabels())),
			PropertyClasses: make(map[string]string, len(cc.GetLabels())),
		}

		for _, p := range lc.GetProperties() {
			values, _ := rxa.GetAllValues(p, id)
			li.PropertyValues[p] = values
		}

		fcp := acp.GetFormatterConfigurator()

		glf := fcp.GetLabelFormatter()
		gcf := fcp.GetClassFormatter()

		for _, l := range cc.GetLabels() {
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
