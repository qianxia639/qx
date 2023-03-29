package qx

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodGet, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPost, pattern, handler)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPut, pattern, handler)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodDelete, pattern, handler)
}

func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr)

	return http.ListenAndServe(address, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
