package main

import (
	"httpWeb/ws"
	"net/http"
)

func main() {
	ws.CreateServer()

	ws.Handle("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("index"))
	})

	// ws.Get("/apple", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("GET " + r.URL.Path))
	// })

	ws.Post("/apple", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST" + r.URL.Path))
	})

	ws.Listen("8080")
}
