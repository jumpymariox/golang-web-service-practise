package ws

import (
	httpmethod "httpWeb/enum"
	"httpWeb/router"
	"log"
	"net/http"
)

// The serverHandler for mux
type serverHandler struct {
	router router.RouterMap
}

// WebServer is core
type WebServer struct {
	mux     *http.ServeMux
	handler *serverHandler
}

// New instance
func New() *WebServer {
	r := router.New()
	handler := &serverHandler{router: r}
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
	return ws.handler.router.Add(httpmethod.GET, pattern, handler)
}

// Post request
func (ws *WebServer) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return ws.handler.router.Add(httpmethod.POST, pattern, handler)
}

// Put request
func (ws *WebServer) Put(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return ws.handler.router.Add(httpmethod.PUT, pattern, handler)
}

// Delete request
func (ws *WebServer) Delete(pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	return ws.handler.router.Add(httpmethod.DELETE, pattern, handler)
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
