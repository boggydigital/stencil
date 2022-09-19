package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func Group(tmpl *template.Template, gvm *view_models.Group, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "group-page", gvm)
}
