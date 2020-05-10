package ws

import (
	"log"
	"net/http"
)

// ServerHandler is ready for http
type ServerHandler struct{}

var mux *http.ServeMux

// CreateServerMux is used to new one server
func CreateServerMux() *http.ServeMux {
	mux = http.NewServeMux()
	mux.Handle("/", &ServerHandler{})
	return mux
}

// Listen port
func Listen(port string) {
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func (s *ServerHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(req.Method + req.URL.Path))
}
