package web

import "net/http"

type Server interface {
	Route(method string, pattern string, handlerFunc func(c *Context))
	Start(address string) error
}

/*
*
Http Server —— 基于 map 的路由
SdkHttpServer这个是基于net/http这个包实现的http server
*/
type SdkHttpServer struct {
	// Name server 的命中，给个标记，日志输出的时候用得上
	Name    string
	handler *HandlerBasedOnMap
}

func (s *SdkHttpServer) Route(method string, pattern string, handlerFunc func(c *Context)) {
	//http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
	//	c := NewContext(w, r)
	//	handlerFunc(c)
	//})
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handlerFunc
}

func (s *SdkHttpServer) Start(address string) error {
	// 每次请求，都需要先经过s.handler进行过滤处理
	return http.ListenAndServe(address, s.handler)
}

func NewSdkHttpServer(name string) Server {
	handlerMap := make(map[string]func(c *Context))
	return &SdkHttpServer{
		Name: name,
		handler: &HandlerBasedOnMap{
			handlers: handlerMap,
		},
	}
}
