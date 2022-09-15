package stencil

import (
	"embed"
	"html/template"
)

var (
	//go:embed "templates/*.gohtml"
	defaultTemplates embed.FS
)

func InitTemplates() *template.Template {
	return template.Must(
		template.New("").
			ParseFS(defaultTemplates, "templates/*.gohtml"))
}
