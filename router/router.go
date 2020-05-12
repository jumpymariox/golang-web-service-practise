package router

import "net/http"

type RouterMap map[string]func(http.ResponseWriter, *http.Request)

func New() RouterMap {
	return make(RouterMap)
}

func (r RouterMap) Add(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) bool {
	key := method + "_" + pattern
	if r[key] != nil {
		return false
	}
	r[key] = handler
	return true
}
