package main

import (
	"httpWeb/ws"
	"net/http"
)

func main() {
	mux := ws.CreateServerMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	ws.Listen("8080")
}
