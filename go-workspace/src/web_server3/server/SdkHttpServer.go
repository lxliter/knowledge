package server

import (
	"net/http"
	"web_server3/web"
)

/**
具体的一种的http server实现
 */
type SdkHttpServer struct {
	Name    string
	handler web.Handler
}

func (s *SdkHttpServer) Route(method string, pattern string, handlerFunc func(c *web.Context)) {
	// 调用handler的注册路由方法
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *SdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s.handler)
}

func NewSdkHttpServer(name string) web.Server {
	// 初始化路由处理器
	handler := &web.RouterHandler{
		Handlers: make(map[string]func(c *web.Context)),
	}
	return &SdkHttpServer{
		Name: name,
		handler: handler,
	}
}
