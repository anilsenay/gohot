package hotreload

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func findAvailablePort() string {
	port := ""
	for port == "" {
		port, _ = tryListen()
	}

	return port
}

func tryListen() (string, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}

	defer listener.Close()
	return strconv.Itoa(listener.Addr().(*net.TCPAddr).Port), nil
}

func isHostAvailable(remote *url.URL) bool {
	_, err := http.DefaultClient.Do(&http.Request{URL: remote})
	return err == nil
}
