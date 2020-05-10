package main

import (
	"httpWeb/ws"
	"net/http"
)

func main() {
	ws.CreateServerMux()

	ws.Handle("/test", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("test"))
	})

	ws.Listen("8080")
}
