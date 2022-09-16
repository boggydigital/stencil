package stencil

import (
	"embed"
	"html/template"
	"io/fs"
)

var (
	//go:embed "templates/*.gohtml"
	defaultTemplates embed.FS
	tmpl             *template.Template
)

func init() {
	tmpl = template.Must(
		template.New("stencil").
			ParseFS(defaultTemplates, "templates/*.gohtml"))
}

func InitAppTemplates(userTemplates fs.FS, pattern string) {
	template.Must(tmpl.ParseFS(userTemplates, pattern))
}
