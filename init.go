package stencil

import (
	"embed"
	"html/template"
	"io/fs"
	"strings"
)

var (
	//go:embed "templates/*.gohtml"
	defaultTemplates embed.FS
	tmpl             *template.Template
)

var templateFuncs = template.FuncMap{
	"join": join,
}

func join(strs []string) string {
	return strings.Join(strs, ", ")
}

func init() {
	tmpl = template.Must(
		template.New("stencil").
			Funcs(templateFuncs).
			ParseFS(defaultTemplates, "templates/*.gohtml"))
}

func InitAppTemplates(userTemplates fs.FS, pattern string) {
	template.Must(tmpl.ParseFS(userTemplates, pattern))
}
