package web

import (
	"fmt"
	"net/http"
)

type HandlerBasedOnMap struct {
	handlers map[string]func(c *Context)
}

func (h *HandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

func (h *HandlerBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 拼接请求方式和路径
	key := h.key(r.Method, r.URL.Path)
	// 判断路由是否存在
	if handler, ok := h.handlers[key]; ok {
		c := NewContext(w, r)
		handler(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("noy any router match"))
	}
}
