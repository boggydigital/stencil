package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func Section(tmpl *template.Template, scvm *view_models.SectionContent, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "iframe-page", scvm)
}
