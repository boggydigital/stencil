package stencil

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil/render"
	"github.com/boggydigital/stencil/view_models"
	"io"
	"net/url"
)

var ScriptHashes = []string{
	//script-iframe-size-receive-message.gohtml
	"'sha256-EoiesIg5jhsIaHn7PSaZ/oT9Yi0MCUx9WzALOyH9HkE='",
	//script-iframe-post-message.gohtml
	"'sha256-vEdzDTUjeRFG21L/pW+qldt1k+gnTSWl4v2E16iqJPc='",
}

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
	searchScopes       []string
	searchScopeQueries map[string]string
	searchProperties   []string
	titleProperty      string
	propertyTitles     map[string]string
	sectionTitles      map[string]string
	digestTitles       map[string]string
}

func NewApp(title, favIconAccent string, rxa kvas.ReduxAssets) *ReduxApp {
	return &ReduxApp{
		page: &view_models.Page{
			Head: &view_models.Header{
				AppTitle:      title,
				FavIconAccent: favIconAccent,
			},
			Foot: &view_models.Footer{},
		},
		rxa: rxa,
	}
}

func (app *ReduxApp) SetFooter(location, repoUrl string) {
	app.page.Foot = &view_models.Footer{
		Location: location,
		RepoUrl:  repoUrl,
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

func (app *ReduxApp) SetSearchParams(
	scopes []string,
	scopeQueries map[string]string,
	properties []string) error {
	app.searchScopes = scopes
	app.searchProperties = properties
	app.searchScopeQueries = make(map[string]string, len(scopeQueries))

	for scope, query := range scopeQueries {
		q, err := url.ParseQuery(query)
		if err != nil {
			return err
		}
		app.searchScopeQueries[scope], err = url.QueryUnescape(q.Encode())
		if err != nil {
			return err
		}
	}
	return nil
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
		view_models.Formatter(app.itemTitleFormatter),
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
		app.searchScopes,
		app.searchScopeQueries,
		query,
		ids,
		app.searchProperties,
		app.titleProperty,
		app.labels,
		app.listProperties,
		app.propertyTitles,
		digests,
		app.digestTitles,
		view_models.Formatter(app.itemTitleFormatter),
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

		app.SetCurrentNav(ivm.Title)

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

func (app *ReduxApp) RenderGroup(
	navItem string,
	groupOrder []string,
	groupItems map[string][]string,
	groupTitles map[string]string,
	updated string,
	w io.Writer) error {

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
		view_models.Formatter(app.itemTitleFormatter),
		updated,
		app.rxa); err != nil {
		return err
	} else {
		if err := render.Group(tmpl, gvm, w); err != nil {
			return err
		}
	}

	return nil
}
