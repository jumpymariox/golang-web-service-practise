package ws

import (
	"log"
	"net/http"
)

// WebServer is core
type WebServer struct {
	mux *http.ServeMux
}

var mux *http.ServeMux

// New instance
func New() *WebServer {
	return &WebServer{}
}

// CreateServer is used to new one server
func (ws *WebServer) CreateServer() {
	ws.mux = http.NewServeMux()
}

// Handle mux
func (ws *WebServer) Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	if ws.mux == nil {
		return false
	}
	ws.mux.HandleFunc(pattern, handler)
	return true
}

// Get request
func (ws *WebServer) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	ws.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}
		handler(w, r)
	})
}

// Post request
func (ws *WebServer) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	ws.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		handler(w, r)
	})
}

// Listen port
func (ws *WebServer) Listen(port string) {
	server := &http.Server{Addr: ":" + port, Handler: ws.mux}
	log.Fatal(server.ListenAndServe())
}
