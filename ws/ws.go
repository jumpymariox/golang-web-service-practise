package ws

import (
	"log"
	"net/http"
)

var mux *http.ServeMux

// CreateServerMux is used to new one server
func CreateServer() *http.ServeMux {
	mux = http.NewServeMux()
	return mux
}

// Handle mux
func Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(pattern, handler)
}

// Get request
func Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}
		handler(w, r)
	})
}

// Post request
func Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		handler(w, r)
	})
}

// Listen port
func Listen(port string) {
	server := &http.Server{Addr: ":" + port, Handler: mux}
	log.Fatal(server.ListenAndServe())
}
