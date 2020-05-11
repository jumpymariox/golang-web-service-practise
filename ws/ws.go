package ws

import (
	httpmethod "httpWeb/enum"
	"log"
	"net/http"
)

type routerMap map[string]func(http.ResponseWriter, *http.Request)

// The serverHandler for mux
type serverHandler struct {
	router routerMap
}

// WebServer is core
type WebServer struct {
	mux     *http.ServeMux
	handler *serverHandler
}

// New instance
func New() *WebServer {
	handler := &serverHandler{router: make(map[string]func(http.ResponseWriter, *http.Request))}
	return &WebServer{mux: http.NewServeMux(), handler: handler}
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
func (ws *WebServer) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return updateRouter(ws.handler.router, httpmethod.GET, pattern, handler)
}

// Post request
func (ws *WebServer) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return updateRouter(ws.handler.router, httpmethod.POST, pattern, handler)
}

// Put request
func (ws *WebServer) Put(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return updateRouter(ws.handler.router, httpmethod.PUT, pattern, handler)
}

// Delete request
func (ws *WebServer) Delete(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return updateRouter(ws.handler.router, httpmethod.DELETE, pattern, handler)
}

// Listen port
func (ws *WebServer) Listen(port string) {
	ws.mux.Handle("/", ws.handler)
	server := &http.Server{Addr: ":" + port, Handler: ws.mux}
	log.Fatal(server.ListenAndServe())
}

func (s *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "_" + r.URL.Path
	handler := s.router[key]
	if handler != nil {
		handler(w, r)
	}
}

func updateRouter(router routerMap, method string, pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	key := method + "_" + pattern
	if router[key] != nil {
		return false
	}
	router[key] = handler
	return true
}
