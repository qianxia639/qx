package qx

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *routr
}

func New() *Engine {
	return &Engine{router: newRoutr()}
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
	c := newContext(w, r)
	engine.router.handle(c)

}
