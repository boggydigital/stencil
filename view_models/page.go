package view_models

import "html/template"

type Page struct {
	Id      string
	Head    *Header
	Nav     *Navigation
	Content template.HTML
	Foot    *Footer
}

func NewPage(acp AppConfigurator, id, content string) *Page {
	page := acp.GetPage()
	page.Id = id
	page.Content = template.HTML(content)
	return page
}
