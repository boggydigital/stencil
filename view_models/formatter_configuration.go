package view_models

type LabelFormatterGetter interface {
	GetLabelFormatter() Formatter
}

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
	LabelFormatterGetter
	TitleFormatterGetter
	HrefFormatterGetter
	ClassFormatterGetter
	ActionFormatterGetter
}

type FormatterConfigurationProviderGetter interface {
	GetFormatterConfigurationProvider() FormatterConfigurationProvider
}
