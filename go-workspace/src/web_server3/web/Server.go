package web

/**
接口组合
定义服务行为接口->可以有多种服务实现方式，当行为都是相同的
 */
type Server interface {
	Routable
	Start(address string) error
}






