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
	listItemPath       string
	listProperties     []string
	itemProperties     []string
	labels             []string
	itemCoverPath      string
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

func (app *ReduxApp) SetListParams(properties []string) error {
	if err := app.rxa.IsSupported(properties...); err != nil {
		return err
	}

	app.listProperties = properties
	return nil
}

func (app *ReduxApp) SetItemParams(
	properties, sections []string) error {
	if err := app.rxa.IsSupported(properties...); err != nil {
		return err
	}

	app.itemProperties = properties
	app.itemSections = sections
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

func (app *ReduxApp) SetLabels(labels []string) error {
	if err := app.rxa.IsSupported(labels...); err != nil {
		return err
	}

	app.labels = labels
	return nil
}

func (app *ReduxApp) SetLinkParams(
	listItemPath, itemCoverPath string,
	fmtTitle, fmtHref PropertyLinkFormatter) {
	app.listItemPath = listItemPath
	app.itemCoverPath = itemCoverPath
	app.itemTitleFormatter = fmtTitle
	app.itemHrefFormatter = fmtHref
}

func (app *ReduxApp) SetCurrentNav(item string) {
	app.page.Nav.Current = item
}

func (app *ReduxApp) RenderList(navItem string, ids []string, w io.Writer) error {

	app.SetCurrentNav(navItem)

	if lvm, err := view_models.NewList(
		app.page,
		app.listItemPath,
		ids,
		app.titleProperty,
		app.labels,
		app.listProperties,
		app.propertyTitles,
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
		app.listItemPath,
		query,
		ids,
		app.searchProperties,
		app.titleProperty,
		app.labels,
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
		app.itemCoverPath,
		app.itemProperties,
		app.titleProperty,
		app.labels,
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

func (app *ReduxApp) RenderSection(id, section, content string, w io.Writer) error {

	cvm := view_models.NewSection(id, section, content)

	if err := render.Section(tmpl, cvm, w); err != nil {
		return err
	}

	return nil
}

func (app *ReduxApp) RenderGroup(navItem string, groupOrder []string, groupItems map[string][]string, groupTitles map[string]string, w io.Writer) error {

	app.SetCurrentNav(navItem)

	if gvm, err := view_models.NewGroup(
		app.page,
		app.listItemPath,
		groupOrder,
		groupItems,
		groupTitles,
		app.titleProperty,
		app.labels,
		app.listProperties,
		app.propertyTitles,
		app.rxa); err != nil {
		return err
	} else {
		if err := render.Group(tmpl, gvm, w); err != nil {
			return err
		}
	}

	return nil
}
