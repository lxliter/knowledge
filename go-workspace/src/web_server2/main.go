package main

import (
	"fmt"
	"web_server2/request"
	"web_server2/response"
	"web_server2/web"
)

/*
*
Http Server —— 支持 RESTFul API
- RESTFul API 定义
- 路由设计 —— Handler 抽象
  - map 语法
  - 基于 map 的 Handler 实现

- 语法：组合
- 重构
简单来说，就是http method 决定了操作，http path 决定了操作对象
Http Server —— 如何支持 RESTFul API
PUT /user 创建用户
POST /user 更新用户
DELETE /user S删除用户
GET /user 获取用户

http method + http path = http handler

Http Server —— 如何支持 RESTFul API
http.HandleFunc 好像不太行，我们得自己做路由了

路由设计 —— Handler 抽象
实现一个 Handler，它负责路由
如果找到了路由，就执行业务代码
找不到就返回 404

Http Server —— 如何路由？
尝试用 map 写一个最简单的版本

Http Server —— 基于 map 的路由
- 和实现 HandlerBasedOnMap强耦合
- Route 方法依赖于知道HandlerBasedOnMap 的内部细节
- 当我们想要扩展到利用路有树来实现的时候，sdkHttpServer也要修改

Http Server —— Handler 抽象
我们给 HandlerBasedOnMap加一个方法：Route
我们希望 sdkHttpServer
依赖于一个接口，所以我们定义一个自己的接口
*/
func main() {
	server := web.NewSdkHttpServer("my-test-webapp")
	server.Route("POST", "/signUp", signUp)
	err := server.Start(":8080")
	if err != nil {
		fmt.Printf("web server start error %v \n", err)
	}
}

/*
*
框架来创建context，就可以完全控制什么时候创建，context可以有什么字段。
作为设计者，这种东西不能交给用户自由发挥。
*/
func signUp(c *web.Context) {
	req := &request.SignUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		_ = c.BadRequestJson(&response.CommonResponse{
			BizCode: 4,
			Msg:     fmt.Sprintf("invalid request: %v", err),
		})
		return
	}
	_ = c.BadRequestJson(&response.CommonResponse{
		// 假设这个是新用户的id
		Data: "6",
	})
}
