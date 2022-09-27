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
	//script_images_fade_in.gohtml
	"'sha256-mHF+LRXfnTXY5tSp3gx9qjKNUv7/AgPvHeyrgmVqJoA='",
}

type App struct {
	page *view_models.Page
	//rxa                kvas.ReduxAssets
	listItemPath       string
	listProperties     []string
	itemProperties     []string
	labels             []string
	icons              []string
	coverPath          string
	itemSections       []string
	itemHrefFormatter  view_models.Formatter
	itemTitleFormatter view_models.Formatter
	searchScopes       []string
	searchScopeQueries map[string]string
	searchProperties   []string
	titleProperty      string
	listCoverProperty  string
	propertyTitles     map[string]string
	sectionTitles      map[string]string
	digestTitles       map[string]string
}

func NewApp(title, favIconAccent string) *App {
	return &App{
		page: &view_models.Page{
			Head: &view_models.Header{
				AppTitle:      title,
				FavIconAccent: favIconAccent,
			},
			Foot: &view_models.Footer{},
		},
	}
}

func (a *App) GetPage() *view_models.Page {
	return a.page
}

func (a *App) GetCoverPath() string {
	return a.coverPath
}

func (a *App) GetLabels() []string {
	return a.labels
}

func (a *App) GetIcons() []string {
	return a.icons
}

func (a *App) GetTitleProperty() string {
	return a.titleProperty
}

func (a *App) GetListItemPath() string {
	return a.listItemPath
}

func (a *App) GetListProperties() []string {
	return a.listProperties
}

func (a *App) GetListCoverProperty() string {
	return a.listCoverProperty
}

func (a *App) GetItemProperties() []string {
	return a.itemProperties
}

func (a *App) GetItemSections() []string {
	return a.itemSections
}

func (a *App) GetItemHrefFormatter() view_models.Formatter {
	return a.itemHrefFormatter
}

func (a *App) GetItemTitleFormatter() view_models.Formatter {
	return a.itemTitleFormatter
}

func (a *App) GetSearchScopes() []string {
	return a.searchScopes
}

func (a *App) GetSearchScopeQueries() map[string]string {
	return a.searchScopeQueries
}

func (a *App) GetSearchProperties() []string {
	return a.searchProperties
}

func (a *App) GetPropertyTitles() map[string]string {
	return a.propertyTitles
}

func (a *App) GetSectionTitles() map[string]string {
	return a.sectionTitles
}

func (a *App) GetDigestTitles() map[string]string {
	return a.digestTitles
}

func (app *App) SetFooter(location, repoUrl string) {
	app.page.Foot = &view_models.Footer{
		Location: location,
		RepoUrl:  repoUrl,
	}
}

func (app *App) SetNavigation(
	items []string,
	icons map[string]string,
	hrefs map[string]string) {
	app.page.Nav = &view_models.Navigation{
		Items: items,
		Icons: icons,
		Hrefs: hrefs,
	}
}

func (app *App) SetListParams(
	listCoverProperty string,
	properties []string,
	rxa kvas.ReduxAssets) error {
	if rxa != nil {
		if err := rxa.IsSupported(properties...); err != nil {
			return err
		}
	}

	app.listCoverProperty = listCoverProperty
	app.listProperties = properties
	return nil
}

func (app *App) SetItemParams(
	properties, sections []string,
	rxa kvas.ReduxAssets) error {
	if rxa != nil {
		if err := rxa.IsSupported(properties...); err != nil {
			return err
		}
	}

	app.itemProperties = properties
	app.itemSections = sections
	return nil
}

func (app *App) SetSearchParams(
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

func (app *App) SetTitles(
	titleProperty string,
	propertyTitles, sectionTitles, digestTitles map[string]string) {
	app.titleProperty = titleProperty
	app.propertyTitles = propertyTitles
	app.sectionTitles = sectionTitles
	app.digestTitles = digestTitles
}

func (app *App) SetLabels(labels []string, rxa kvas.ReduxAssets) error {
	if rxa != nil {
		if err := rxa.IsSupported(labels...); err != nil {
			return err
		}
	}

	app.labels = labels
	return nil
}

func (app *App) SetIcons(icons []string, rxa kvas.ReduxAssets) error {
	if rxa != nil {
		if err := rxa.IsSupported(icons...); err != nil {
			return err
		}
	}

	app.icons = icons
	return nil
}

func (app *App) SetLinkParams(
	listItemPath, coverPath string,
	fmtTitle, fmtHref view_models.Formatter) {
	app.listItemPath = listItemPath
	app.coverPath = coverPath
	app.itemTitleFormatter = fmtTitle
	app.itemHrefFormatter = fmtHref
}

func (app *App) SetCurrentNav(item string) {
	app.page.Nav.Current = item
}

func (app *App) RenderList(navItem string, ids []string, rxa kvas.ReduxAssets, w io.Writer) error {

	app.SetCurrentNav(navItem)

	if lvm, err := view_models.NewList(app, ids, rxa); err != nil {
		return err
	} else {
		if err := render.List(tmpl, lvm, w); err != nil {
			return err
		}
	}
	return nil
}

func (app *App) RenderSearch(
	navItem string,
	query map[string][]string,
	ids []string,
	digests map[string][]string,
	rxa kvas.ReduxAssets,
	w io.Writer) error {
	app.SetCurrentNav(navItem)

	if svm, err := view_models.NewSearch(
		app,
		query,
		ids,
		digests,
		rxa); err != nil {
		return err
	} else {
		if err := render.Search(tmpl, svm, w); err != nil {
			return err
		}
	}
	return nil
}

func (app *App) RenderItem(id string, rxa kvas.ReduxAssets, w io.Writer) error {

	if ivm, err := view_models.NewItem(app, id, rxa); err != nil {
		return err
	} else {

		app.SetCurrentNav(ivm.Title)

		if err := render.Item(tmpl, ivm, w); err != nil {
			return err
		}
	}
	return nil
}

func (app *App) RenderSection(id, section, content string, w io.Writer) error {

	cvm := view_models.NewSection(id, section, content)

	if err := render.Section(tmpl, cvm, w); err != nil {
		return err
	}

	return nil
}

func (app *App) RenderGroup(
	navItem string,
	groupOrder []string,
	groupItems map[string][]string,
	groupTitles map[string]string,
	updated string,
	rxa kvas.ReduxAssets,
	w io.Writer) error {

	app.SetCurrentNav(navItem)

	if gvm, err := view_models.NewGroup(
		app,
		groupOrder,
		groupItems,
		groupTitles,
		updated,
		rxa); err != nil {
		return err
	} else {
		if err := render.Group(tmpl, gvm, w); err != nil {
			return err
		}
	}

	return nil
}
