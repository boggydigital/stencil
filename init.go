package stencil

import (
	"embed"
	"github.com/boggydigital/issa"
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
	"commaJoin": commaJoin,
	"concat":    concat,
	"contains":  contains,
	"hydrate":   hydrate,
}

func contains(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func commaJoin(strs []string) string {
	return strings.Join(strs, ", ")
}

func concat(strs ...string) string {
	return strings.Join(strs, "")
}

func hydrate(dhi string) template.URL {
	return template.URL(issa.HydrateColor(dhi))
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
