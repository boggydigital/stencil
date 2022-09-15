package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func Item(tmpl *template.Template, ivm *view_models.Item, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "item-page", ivm)
}
