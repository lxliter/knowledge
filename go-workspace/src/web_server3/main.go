package main

import (
	"fmt"
	"web_server3/request"
	"web_server3/response"
	"web_server3/server"
	"web_server3/web"
)

func main() {
	server := server.NewSdkHttpServer("my-test-webapp")
	server.Route("POST", "/signUp", signUp)
	err := server.Start(":8080")
	if err != nil {
		fmt.Printf("web server start error %v \n", err)
	}
}

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
		Data: "666666",
	})
}
