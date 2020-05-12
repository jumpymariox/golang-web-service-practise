package main

import (
	"httpWeb/ws"
	"net/http"
)

func main() {
	w := ws.New()

	w.Get("/apple", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET " + r.URL.Path))
	})

	w.Post("/apple", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST " + r.URL.Path))
	})

	w.Listen("8080")
}
