package stencil

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil/render"
	"github.com/boggydigital/stencil/view_models"
	"io"
)

type ReduxApp struct {
	Title          string
	FavIcon        string
	rxa            kvas.ReduxAssets
	nav            *view_models.Navigation
	listProperties []string
	listItemHref   string
}

func NewApp(title, favIcon string, rxa kvas.ReduxAssets) *ReduxApp {
	return &ReduxApp{
		Title:   title,
		FavIcon: favIcon,
		rxa:     rxa,
	}
}

func (app *ReduxApp) SetNavigation(
	items []string,
	icons map[string]string,
	hrefs map[string]string) {
	app.nav = &view_models.Navigation{
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

func (app *ReduxApp) RenderList(navItem string, ids []string, w io.Writer) error {

	if navItem != "" {
		app.nav.Current = navItem
	}

	if lvm, err := view_models.NewList(
		app.nav,
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
