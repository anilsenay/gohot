package hotreload

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type transport struct {
	http.RoundTripper
	addScripts func(b []byte) []byte
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if strings.Contains(req.Header.Get("Accept"), "text/html") && bytes.Contains(b, []byte("<html>")) {
		b = t.addScripts(b)
	}

	body := io.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return resp, nil
}
