package qx

import "net/http"

type routr struct {
	handlers map[string]HandlerFunc
}

func newRoutr() *routr {
	return &routr{handlers: make(map[string]HandlerFunc)}
}

func (r *routr) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "" + pattern
	r.handlers[key] = handler
}
func (r *routr) handle(c *Context) {
	key := c.Request.Method + "" + c.Request.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.Status(http.StatusNotFound)
	}
}
