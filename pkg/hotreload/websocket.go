package hotreload

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func webSocket(port string, refreshCh <-chan struct{}) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		HandshakeTimeout: time.Second * 10,
	}

	connPool := map[string]*websocket.Conn{}

	mux := http.NewServeMux()

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		connPool[conn.RemoteAddr().String()] = conn

		defer conn.Close()

		for range refreshCh {
			for addr, c := range connPool {
				err := c.WriteMessage(websocket.TextMessage, []byte("refresh"))
				if err != nil {
					c.Close()
					delete(connPool, addr)
				}
			}
		}
	})

	err := http.ListenAndServe(":"+port, mux)
	fmt.Println(err)
}
