package stencil_rest

import (
	"net/http"
	"os"
	"path/filepath"
)

const (
	relRobotsTxtFilename = "robots.txt"
)

func GetRobotsTxt(dir string, w http.ResponseWriter, r *http.Request) {
	absRobotsTxt := filepath.Join(dir, relRobotsTxtFilename)
	if _, err := os.Stat(absRobotsTxt); os.IsNotExist(err) {
		http.NotFound(w, r)
	} else {
		http.ServeFile(w, r, absRobotsTxt)
	}
}
