package web

import (
	"fmt"
	"net/http"
)

var _ Handler = &RouterHandler{}

/*
*
路由处理器，负责路由的注册以及拦截
*/
type RouterHandler struct {
	Handlers map[string]func(c *Context)
}

func (rh *RouterHandler) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

func (rh *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := rh.key(r.Method, r.URL.Path)
	if handler, ok := rh.Handlers[key]; ok {
		c := NewContext(w, r)
		handler(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("noy any router match"))
	}
}

func (rh *RouterHandler) Route(method string, pattern string, handlerFunc func(c *Context)) {
	routeKey := rh.key(method, pattern)
	rh.Handlers[routeKey] = handlerFunc
}
