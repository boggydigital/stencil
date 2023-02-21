package view_models

type PropertyEditor struct {
	*Page
	Id                 string
	Title              string
	PropertyTitle      string
	Condition          bool
	ConditionalMessage string
	SelectedValues     map[string]bool
	AllValues          []string
	AllowNewValues     bool
	ApplyEndpoint      string
}

func NewPropertyEditor(
	app AppConfigurator,
	id, title, propertyTitle string,
	condition bool, conditionalMessage string,
	selectedValues map[string]bool, allValues []string, allowNewValues bool,
	applyEndpoint string) *PropertyEditor {
	return &PropertyEditor{
		Page:               app.GetPage(),
		Id:                 id,
		Title:              title,
		PropertyTitle:      propertyTitle,
		Condition:          condition,
		ConditionalMessage: conditionalMessage,
		SelectedValues:     selectedValues,
		AllValues:          allValues,
		AllowNewValues:     allowNewValues,
		ApplyEndpoint:      applyEndpoint,
	}
}
