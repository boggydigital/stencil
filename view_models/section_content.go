package view_models

import "html/template"

type SectionContent struct {
	Id      string
	Content template.HTML
}

func NewSectionContent(id, desc string) *SectionContent {
	return &SectionContent{
		Id:      id,
		Content: template.HTML(desc),
	}
}
