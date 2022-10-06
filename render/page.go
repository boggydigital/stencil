package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func Page(tmpl *template.Template, pvm *view_models.Page, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "page", pvm)
}
