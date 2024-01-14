package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Print("Backend running on http://localhost:8080")
	log.Print("Frontend running on http://localhost:3000")
	_ = http.ListenAndServe(":8080", nil)
}
