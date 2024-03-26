package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("web server start error %v", err)
	}
}

/**
HttpServer——Server和Context
HttpServer实现
Context抽象与实现
- 读取数据
- 写入响应
创建Context
 */

/**
要点总结：type
type定义熟记。其中typeA=B这种别名，一般只用于兼容性处理，所以不需要过多关注；
先有抽象再有实现，所以要先定义接口
鸭子类型：一个结构体有某个接口的所有方法，它就实现了这个接口；
指针：方法接收器，遇事不决用指针；
 */

/**
基础语法——struct定义
基本语法:
type Name struct{
	fieldName FieldType
    //...
}
结构体和结构体的字段都遵循大小写控制访问性的原则

Tip：***其实还有别的第三方http库，也可以用来实现一个server***
*/

// sdkHttpServer这个是基于net/http这个包实现的http server
type sdkHttpServer struct {
	// Name server的名字，给个标记，日志输出的时候用得上
	Name string
}

/**
当一个结构体具备接口的所有的方法时候，它就实现了这个接口
 */

// Route 设定一个路由，命中该路由的会执行handlerFunc的代码
func (s sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

// Start 启动我们的服务器
func (s sdkHttpServer) Start(address string) error {
	//TODO implement me
	panic("implement me")
}

/*
*
HttpServer抽象——接口定义
基础语法——interface定义
基本语法type名字interface{}
里面只能有方法，方法也不需要func关键字
啥是接口（interface）：接口是一组行为的抽象
尽量用接口，以实现面向接口编程

Tip：当你怀疑要不要用接口的时候，加上去总是很保险的

注释规范：以被注释的开头，后面跟着描述
*/

// Server是http server的顶级抽象
type Server interface {
	// Route 设定一个路由，命中该路由的会执行handlerFunc的代码
	Route(pattern string, handlerFunc http.HandlerFunc)

	// Start 启动我们的服务器
	Start(address string) error
}

/*
*
从HttpServer开始
如果我想启动两个服务器，一个监听8080，一个监听8081，用8081来作为管理端口
这个东西，缺乏一个逻辑上的联系，至少联系不够紧密

HttpServer抽象
我想要一个Server的东西，***表达一种逻辑上的抽象***，
它代表的是对某个端口的进行监听的实体，必要的时候，我可以开启多个Server，来监听多个端口
*/
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is home page")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is login page")
}

/**
基础语法——type定义
type定义
	type名字interface{}
	type名字struct{}
	type 名字 别的类型
	type 别名 = 别的类型
结构体初始化
指针与方法接收器
结构体如何实现接口
*/
