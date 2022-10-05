package view_models

type TitleFormatterGetter interface {
	GetTitleFormatter() Formatter
}

type HrefFormatterGetter interface {
	GetHrefFormatter() Formatter
}

type ClassFormatterGetter interface {
	GetClassFormatter() Formatter
}

type ActionFormatterGetter interface {
	GetActionFormatter() Formatter
}

type FormatterConfigurationProvider interface {
	TitleFormatterGetter
	HrefFormatterGetter
	ClassFormatterGetter
	ActionFormatterGetter
}

type FormatterConfigurationProviderGetter interface {
	GetFormatterConfigurationProvider() FormatterConfigurationProvider
}
