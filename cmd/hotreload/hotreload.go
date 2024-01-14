package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/anilsenay/go-hot-reload/pkg/hotreload"
)

func main() {
	var mode string
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	port := flag.String("port", "3000", "Port to use hot reloading server")
	proxy := flag.String("proxy", "http://localhost:8080", "Proxy address of an existing server")

	switch mode {
	case "start":
		startHotReloadServer(*port, *proxy)
	case "refresh":
		refreshServer(*port)
	default:
		startHotReloadServer(*port, *proxy)
	}
}

func startHotReloadServer(port string, proxyAddr string) {
	hotReload := hotreload.New(port, proxyAddr)
	_ = hotReload.Listen()
}

func refreshServer(port string) {
	_, err := http.Get("http://localhost:" + port + "/refresh")
	if err != nil {
		fmt.Println("request error: ", err.Error())
	}
}
