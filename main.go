package main

import (
	"log"
	"net/http"
)

// The ServeHandler interface
type ServeHandler struct{}

func (s *ServeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("start"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &ServeHandler{})
	mux.HandleFunc("/hello", sayHello)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func sayHello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello"))
}
