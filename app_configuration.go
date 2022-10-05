package stencil

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil/render"
	"github.com/boggydigital/stencil/view_models"
	"io"
	"net/url"
)

type AppConfiguration struct {
	page         *view_models.Page
	commonConfig *CommonConfiguration
	listConfig   *ListConfiguration
	itemConfig   *ItemConfiguration
	fmtConfig    *FormatterConfiguration
	searchConfig *SearchConfiguration
}

func NewAppConfig(title, favIconAccent string) *AppConfiguration {
	return &AppConfiguration{
		page: &view_models.Page{
			Head: &view_models.Header{
				AppTitle:      title,
				FavIconAccent: favIconAccent,
			},
			Foot: &view_models.Footer{},
		},
		commonConfig: &CommonConfiguration{},
		listConfig:   &ListConfiguration{},
		itemConfig:   &ItemConfiguration{},
		fmtConfig:    &FormatterConfiguration{},
		searchConfig: &SearchConfiguration{},
	}
}

func (a *AppConfiguration) GetPage() *view_models.Page {
	return a.page
}

func (a *AppConfiguration) GetListConfigurationProvider() view_models.ListConfigurationProvider {
	return a.listConfig
}

func (a *AppConfiguration) GetItemConfigurationProvider() view_models.ItemConfigurationProvider {
	return a.itemConfig
}

func (a *AppConfiguration) GetFormatterConfigurationProvider() view_models.FormatterConfigurationProvider {
	return a.fmtConfig
}

func (a *AppConfiguration) GetSearchConfigurationProvider() view_models.SearchConfigurationProvider {
	return a.searchConfig
}

func (a *AppConfiguration) GetCommonConfigurationProvider() view_models.CommonConfigurationProvider {
	return a.commonConfig
}

func (a *AppConfiguration) SetFooter(location, repoUrl string) {
	a.page.Foot = &view_models.Footer{
		Location: location,
		RepoUrl:  repoUrl,
	}
}

func (a *AppConfiguration) SetNavigation(
	items []string,
	icons map[string]string,
	hrefs map[string]string) {
	a.page.Nav = &view_models.Navigation{
		Items: items,
		Icons: icons,
		Hrefs: hrefs,
	}
}

func (a *AppConfiguration) SetListConfiguration(
	properties, hiddenProperties []string,
	itemPath, imageProperty, imagePath string,
	rxa kvas.ReduxAssets) error {
	if rxa != nil {
		if err := rxa.IsSupported(properties...); err != nil {
			return err
		}
	}

	a.listConfig.properties = properties
	a.listConfig.itemPath = itemPath
	a.listConfig.imageProperty = imageProperty
	a.listConfig.imagePath = imagePath
	a.listConfig.hiddenProperties = hiddenProperties
	return nil
}

func (a *AppConfiguration) SetItemConfiguration(
	properties, computedProperties, hiddenProperties []string,
	sections []string,
	imageProperty, imagePath string,
	rxa kvas.ReduxAssets) error {
	if rxa != nil {
		if err := rxa.IsSupported(properties...); err != nil {
			return err
		}
	}

	a.itemConfig.properties = properties
	a.itemConfig.computedProperties = computedProperties
	a.itemConfig.hiddenProperties = hiddenProperties
	a.itemConfig.sections = sections
	a.itemConfig.imageProperty = imageProperty
	a.itemConfig.imagePath = imagePath
	return nil
}

func (a *AppConfiguration) SetFormatterConfiguration(
	fmtLabel, fmtTitle, fmtHref, fmtClass, fmtAction view_models.Formatter) {

	a.fmtConfig.labelFormatter = fmtLabel
	a.fmtConfig.titleFormatter = fmtTitle
	a.fmtConfig.hrefFormatter = fmtHref
	a.fmtConfig.classFormatter = fmtClass
	a.fmtConfig.actionFormatter = fmtAction
}

func (a *AppConfiguration) SetSearchConfiguration(
	properties []string,
	scopes []string,
	scopeQueries map[string]string) error {

	a.searchConfig.properties = properties
	a.searchConfig.scopes = scopes
	a.searchConfig.scopeQueries = make(map[string]string, len(scopeQueries))

	for scope, query := range scopeQueries {
		q, err := url.ParseQuery(query)
		if err != nil {
			return err
		}
		a.searchConfig.scopeQueries[scope], err = url.QueryUnescape(q.Encode())
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AppConfiguration) SetCommonConfiguration(
	labels, hiddenLabels, icons []string,
	titleProperty string,
	propertyTitles, sectionTitles, digestTitles map[string]string,
	rxa kvas.ReduxAssets) error {

	if rxa != nil {
		if err := rxa.IsSupported(append(labels, icons...)...); err != nil {
			return err
		}
	}

	a.commonConfig.labels = labels
	a.commonConfig.hiddenLabels = hiddenLabels
	a.commonConfig.icons = icons
	a.commonConfig.titleProperty = titleProperty
	a.commonConfig.propertyTitles = propertyTitles
	a.commonConfig.sectionTitles = sectionTitles
	a.commonConfig.digestTitles = digestTitles

	return nil
}

func (a *AppConfiguration) SetCurrentNav(item string) {
	a.page.Nav.Current = item
}

func (a *AppConfiguration) RenderList(navItem string, ids []string, rxa kvas.ReduxAssets, w io.Writer) error {

	a.SetCurrentNav(navItem)

	if lvm, err := view_models.NewList(a, ids, rxa); err != nil {
		return err
	} else {
		if err := render.List(tmpl, lvm, w); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppConfiguration) RenderSearch(
	navItem string,
	query map[string][]string,
	ids []string,
	digests map[string][]string,
	rxa kvas.ReduxAssets,
	w io.Writer) error {

	a.SetCurrentNav(navItem)

	if svm, err := view_models.NewSearch(
		a,
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

func (a *AppConfiguration) RenderItem(id string, hasSections []string, rxa kvas.ReduxAssets, w io.Writer) error {

	if ivm, err := view_models.NewItem(a, id, hasSections, rxa); err != nil {
		return err
	} else {

		a.SetCurrentNav(ivm.Title)

		if err := render.Item(tmpl, ivm, w); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppConfiguration) RenderSection(id, section, content string, w io.Writer) error {

	cvm := view_models.NewSection(id, section, content)

	if err := render.Section(tmpl, cvm, w); err != nil {
		return err
	}

	return nil
}

func (a *AppConfiguration) RenderGroup(
	navItem string,
	groupOrder []string,
	groupItems map[string][]string,
	groupTitles map[string]string,
	updated string,
	rxa kvas.ReduxAssets,
	w io.Writer) error {

	a.SetCurrentNav(navItem)

	if gvm, err := view_models.NewGroup(
		a,
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
