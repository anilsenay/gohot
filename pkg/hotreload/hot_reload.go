package hotreload

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type HotReload struct {
	mux       *http.ServeMux
	refreshCh chan struct{}
	port      string
	wsPort    string
}

func New(port string, target string) *HotReload {
	mux := http.NewServeMux()
	refreshCh := make(chan struct{})

	remote, err := url.Parse(target)
	if err != nil {
		panic("Proxy address parse error: " + err.Error())
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	wsPort := findAvailablePort()

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Transport = &transport{
		RoundTripper: http.DefaultTransport,
		addScripts:   generate(wsPort),
	}

	mux.HandleFunc("/", handler(proxy))
	mux.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			for {
				select {
				case <-time.After(time.Second * 10):
					return
				default:
					if isHostAvailable(remote) {
						refreshCh <- struct{}{}
						return
					}
				}
			}
		}()
	})

	return &HotReload{
		mux:       mux,
		refreshCh: refreshCh,
		port:      port,
		wsPort:    wsPort,
	}
}

func (hl *HotReload) Listen() error {
	go webSocket(hl.wsPort, hl.refreshCh)
	return http.ListenAndServe(":"+hl.port, hl.mux)
}
