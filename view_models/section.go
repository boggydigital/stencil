package view_models

import "html/template"

type Section struct {
	Id      string
	Section string
	Content template.HTML
}

func NewSection(id, section, content string) *Section {
	return &Section{
		Id:      id,
		Section: section,
		Content: template.HTML(content),
	}
}
