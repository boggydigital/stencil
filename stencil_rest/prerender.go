package stencil_rest

import (
	"bytes"
	"fmt"
	"github.com/boggydigital/middleware"
	"io"
	"net/http"
)

func Prerender(paths []string, clearStaticContent bool, port int) bool {

	// we don't want to accumulate existing static content over the lifetime of the app
	if clearStaticContent {
		middleware.ClearStaticContent()
	}

	host := fmt.Sprintf("http://localhost:%d", port)

	for _, p := range paths {
		if err := setStaticContent(host, p); err != nil {
			return false
		}
	}

	return true
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
