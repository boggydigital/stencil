package view_models

type PageGetter interface {
	GetPage() *Page
}

type CoverPathGetter interface {
	GetCoverPath() string
}

type LabelsGetter interface {
	GetLabels() []string
}

type IconsGetter interface {
	GetIcons() []string
}

type TitlePropertyGetter interface {
	GetTitleProperty() string
}

type ListItemPathGetter interface {
	GetListItemPath() string
}

type ListPropertiesGetter interface {
	GetListProperties() []string
}

type ListCoverPropertyGetter interface {
	GetListCoverProperty() string
}

type ListClassPropertiesGetter interface {
	GetListClassProperties() []string
}

type ItemPropertiesGetter interface {
	GetItemProperties() []string
}

type ItemSectionsGetter interface {
	GetItemSections() []string
}

type ItemHrefFormatterGetter interface {
	GetItemHrefFormatter() Formatter
}

type ItemTitleFormatterGetter interface {
	GetItemTitleFormatter() Formatter
}

type ItemClassFormatterGetter interface {
	GetItemClassFormatter() Formatter
}

type SearchScopesGetter interface {
	GetSearchScopes() []string
}

type SearchScopeQueriesGetter interface {
	GetSearchScopeQueries() map[string]string
}

type SearchPropertiesGetter interface {
	GetSearchProperties() []string
}

type PropertyTitlesGetter interface {
	GetPropertyTitles() map[string]string
}

type SectionTitlesGetter interface {
	GetSectionTitles() map[string]string
}

type DigestTitlesGetter interface {
	GetDigestTitles() map[string]string
}

type AppConfigurationProvider interface {
	/* Common */
	PageGetter
	CoverPathGetter
	LabelsGetter
	IconsGetter
	TitlePropertyGetter
	/* List */
	ListItemPathGetter
	ListPropertiesGetter
	ListCoverPropertyGetter
	ListClassPropertiesGetter
	/* Item */
	ItemPropertiesGetter
	ItemSectionsGetter
	ItemHrefFormatterGetter
	ItemTitleFormatterGetter
	ItemClassFormatterGetter
	/* Search */
	SearchScopesGetter
	SearchScopeQueriesGetter
	SearchPropertiesGetter
	/* Titles */
	PropertyTitlesGetter
	SectionTitlesGetter
	DigestTitlesGetter
}
