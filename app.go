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
	itemHrefFormatter  PropertyLinkFormatter
	itemTitleFormatter PropertyLinkFormatter
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
	fmtTitle, fmtHref PropertyLinkFormatter) {
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

	if lvm, err := view_models.NewList(
		app.page,
		app.listItemPath,
		ids,
		app.titleProperty,
		app.listCoverProperty,
		app.coverPath,
		app.labels,
		app.icons,
		app.listProperties,
		app.propertyTitles,
		view_models.Formatter(app.itemTitleFormatter),
		rxa); err != nil {
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
		app.page,
		app.listItemPath,
		app.searchScopes,
		app.searchScopeQueries,
		query,
		ids,
		app.searchProperties,
		app.titleProperty,
		app.listCoverProperty,
		app.coverPath,
		app.labels,
		app.icons,
		app.listProperties,
		app.propertyTitles,
		digests,
		app.digestTitles,
		view_models.Formatter(app.itemTitleFormatter),
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

	if ivm, err := view_models.NewItem(
		app.page,
		id,
		app.coverPath,
		app.itemProperties,
		app.titleProperty,
		app.labels,
		app.propertyTitles,
		app.itemSections,
		app.sectionTitles,
		view_models.Formatter(app.itemTitleFormatter),
		view_models.Formatter(app.itemHrefFormatter),
		rxa); err != nil {
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
		app.page,
		app.listItemPath,
		groupOrder,
		groupItems,
		groupTitles,
		app.titleProperty,
		app.listCoverProperty,
		app.coverPath,
		app.labels,
		app.icons,
		app.listProperties,
		app.propertyTitles,
		view_models.Formatter(app.itemTitleFormatter),
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
