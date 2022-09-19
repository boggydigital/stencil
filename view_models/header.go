package view_models

import "html/template"

type Header struct {
	AppTitle string
	FavIcon  template.HTML
}
