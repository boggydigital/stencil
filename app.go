package stencil

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil/render"
	"github.com/boggydigital/stencil/view_models"
	"io"
)

type ReduxApp struct {
	page             *view_models.Page
	rxa              kvas.ReduxAssets
	listItemHref     string
	listProperties   []string
	searchProperties []string
	propertyTitles   map[string]string
	digestTitles     map[string]string
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

func (app *ReduxApp) SetListParams(itemHref string, properties []string) error {
	if err := app.rxa.IsSupported(properties...); err != nil {
		return err
	}

	app.listItemHref = itemHref
	app.listProperties = properties
	return nil
}

func (app *ReduxApp) SetSearchParams(properties []string) {
	app.searchProperties = properties
}

func (app *ReduxApp) SetTitles(propertyTitles map[string]string, digestTitles map[string]string) {
	app.propertyTitles = propertyTitles
	app.digestTitles = digestTitles
}

func (app *ReduxApp) SetCurrentNav(item string) {
	if item != "" {
		app.page.Nav.Current = item
	}
}

func (app *ReduxApp) RenderList(navItem string, ids []string, w io.Writer) error {

	app.SetCurrentNav(navItem)

	if lvm, err := view_models.NewList(
		app.page,
		app.listItemHref,
		ids,
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
