package render

import (
	"github.com/boggydigital/stencil/view_models"
	"html/template"
	"io"
)

func Search(tmpl *template.Template, svm *view_models.Search, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "search-page", svm)
}
