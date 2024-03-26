package web

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type SdkHttpServer struct {
	Name string
}

func (s *SdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s *SdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewSdkHttpServer(name string) Server {
	return &SdkHttpServer{
		Name: name,
	}
}
