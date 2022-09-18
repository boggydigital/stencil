package stencil

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil/render"
	"github.com/boggydigital/stencil/view_models"
	"io"
)

type PropertyLinkFormatter view_models.Formatter

type ReduxApp struct {
	page               *view_models.Page
	rxa                kvas.ReduxAssets
	listItemHref       string
	listProperties     []string
	listLabels         []string
	itemProperties     []string
	itemLabels         []string
	itemCoverHref      string
	itemSections       []string
	itemHrefFormatter  PropertyLinkFormatter
	itemTitleFormatter PropertyLinkFormatter
	searchProperties   []string
	titleProperty      string
	propertyTitles     map[string]string
	sectionTitles      map[string]string
	digestTitles       map[string]string
}

func NewApp(title, favIcon string, rxa kvas.ReduxAssets) *ReduxApp {
	return &ReduxApp{
		page: &view_models.Page{
			Head: &view_models.Header{
				Title:   title,
				FavIcon: favIcon,
			},
		},
		rxa: rxa,
	}
}

func (app *ReduxApp) SetNavigation(
	items []string,
	icons map[string]string,
	hrefs map[string]string) {
	app.page.Nav = &view_models.Navigation{
		Items: items,
		Icons: icons,
		Hrefs: hrefs,
	}
}

func (app *ReduxApp) SetListParams(
	itemHref string,
	properties,
	labels []string) error {
	if err := app.rxa.IsSupported(properties...); err != nil {
		return err
	}

	app.listItemHref = itemHref
	app.listProperties = properties
	return nil
}

func (app *ReduxApp) SetItemParams(
	coverHref string,
	properties, labels, sections []string,
	fmtTitle, fmtHref PropertyLinkFormatter) error {
	if err := app.rxa.IsSupported(properties...); err != nil {
		return err
	}

	app.itemCoverHref = coverHref
	app.itemProperties = properties
	app.itemLabels = labels
	app.itemSections = sections
	app.itemTitleFormatter = fmtTitle
	app.itemHrefFormatter = fmtHref
	return nil
}

func (app *ReduxApp) SetSearchParams(properties []string) {
	app.searchProperties = properties
}

func (app *ReduxApp) SetTitles(
	titleProperty string,
	propertyTitles, sectionTitles, digestTitles map[string]string) {
	app.titleProperty = titleProperty
	app.propertyTitles = propertyTitles
	app.sectionTitles = sectionTitles
	app.digestTitles = digestTitles
}

func (app *ReduxApp) SetCurrentNav(item string) {
	app.page.Nav.Current = item
}

func (app *ReduxApp) RenderList(navItem string, ids []string, w io.Writer) error {

	app.SetCurrentNav(navItem)

	if lvm, err := view_models.NewList(
		app.page,
		app.listItemHref,
		ids,
		app.titleProperty,
		app.listProperties,
		app.rxa); err != nil {
		return err
	} else {
		if err := render.List(tmpl, lvm, w); err != nil {
			return err
		}
	}
	return nil
}

func (app *ReduxApp) RenderSearch(
	navItem string,
	query map[string][]string,
	ids []string,
	digests map[string][]string,
	w io.Writer) error {
	app.SetCurrentNav(navItem)

	if svm, err := view_models.NewSearch(
		app.page,
		app.listItemHref,
		query,
		ids,
		app.searchProperties,
		app.titleProperty,
		app.listProperties,
		app.propertyTitles,
		digests,
		app.digestTitles,
		app.rxa); err != nil {
		return err
	} else {
		if err := render.Search(tmpl, svm, w); err != nil {
			return err
		}
	}
	return nil
}

func (app *ReduxApp) RenderItem(id string, w io.Writer) error {

	app.SetCurrentNav("")

	if ivm, err := view_models.NewItem(
		app.page,
		id,
		app.itemCoverHref,
		app.itemProperties,
		app.titleProperty,
		app.propertyTitles,
		app.itemSections,
		app.sectionTitles,
		view_models.Formatter(app.itemTitleFormatter),
		view_models.Formatter(app.itemHrefFormatter),
		app.rxa); err != nil {
		return err
	} else {
		if err := render.Item(tmpl, ivm, w); err != nil {
			return err
		}
	}
	return nil
}
