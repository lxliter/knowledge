package web

/**
当发现多个结构体具有相同方法时，就需要考虑定义接口了
 */
type Routable interface {
	Route(method string, pattern string, handlerFunc func(c *Context))
}
