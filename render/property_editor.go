package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func PropertyEditor(tmpl *template.Template, pevm *view_models.PropertyEditor, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "property-editor-page", pevm)
}
