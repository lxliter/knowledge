package web

import (
	"net/http"
)

/**
处理路由
1、组合http自带的路由过滤器对应的接口
2、注册路由
 */
type Handler interface {
	http.Handler
	Routable
}


