package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func List(tmpl *template.Template, lvm *view_models.List, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "list-page", lvm)
}
