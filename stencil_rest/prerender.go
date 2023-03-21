package stencil_rest

import (
	"bytes"
	"fmt"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"io"
	"net/http"
)

func Prerender(paths []string, port int, w http.ResponseWriter) {

	// we don't want to accumulate existing static content over the lifetime of the app
	middleware.ClearStaticContent()

	host := fmt.Sprintf("http://localhost:%d", port)

	for _, p := range paths {
		if err := setStaticContent(host, p); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	if _, err := io.WriteString(w, "ok"); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}

func setStaticContent(host, p string) error {
	resp, err := http.DefaultClient.Get(host + p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs := make([]byte, 0, 1024*1024)
	bb := bytes.NewBuffer(bs)

	if _, err := io.Copy(bb, resp.Body); err != nil {
		return err
	}

	middleware.SetStaticContent(p, bb.Bytes())

	return nil
}
