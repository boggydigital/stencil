package view_models

type Navigation struct {
	Items   []string
	Icons   map[string]string
	Hrefs   map[string]string
	Current string
}
