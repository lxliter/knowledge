package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"web_server/request"
	"web_server/response"
	"web_server/web"
)

/*
*
最简单的web服务器——官网例子
Tip：要熟练掌握找官网，找靠谱例子的能力

最简单的web服务器
•直接Idea启动main函数
•浏览器输入http://localhost:8080/golang
*/
func main() {
	server := web.NewSdkHttpServer("my-test-webapp")
	server.Route("/", home)
	server.Route("/order", order)
	server.Route("/order/create", createOrder)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/signUp", signUp)
	server.Route("/signUp2", signUp2)
	server.Route("/signUp3", signUp3)
	server.Route("/signUp4", signUp4)

	err := server.Start(":8080")
	if err != nil {
		fmt.Printf("web server start error %v \n", err)
	}
}

/**
Http Context—— 让 web 框架来创建 context
框架来创建context，就可以完全控制什么时候创建，context可以有什么字段。
作为设计者，这种东西不能交给用户自由发挥。
原本的 Router 方法已经不行了，需要改造
 */

func signUp4(w http.ResponseWriter, r *http.Request) {
	c := web.NewContext(w, r)
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
		Data: "1",
	})
}

func signUp3(w http.ResponseWriter, r *http.Request) {
	c := web.NewContext(w, r)
	req := &request.SignUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		resp := &response.CommonResponse{
			BizCode: 4, // 假如说我们这个代表输入参数错误
			Msg:     fmt.Sprintf("invalid request: %v", err),
		}
		// 转成json格式
		respBytes, _ := json.Marshal(resp)
		fmt.Fprint(w, string(respBytes))
		return
	}
	resp := &response.CommonResponse{
		BizCode: 1, // 代表处理成功
		Msg:     fmt.Sprintf("request data: %v", req),
	}
	// 转成json格式
	respBytes, _ := json.Marshal(resp)
	fmt.Fprint(w, string(respBytes))
}

/*
*
Http Server —— 写入响应
- 难以输出格式化数据，比如说返回一个 json 数据给客户端
- 没有处理 http 响应码
*/
func signUp2(w http.ResponseWriter, r *http.Request) {
	c := web.NewContext(w, r)
	req := &request.SignUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "invalid request: %v", err)
	}
	fmt.Fprintf(w, "request data: %v", req)
}

/*
*
基础语法 —— 空接口 interface{}
- 空接口 interface{} 不包含任何方法
- 所以任何结构体都实现了该接口
- 类似于 Java 的 Object， 即所谓的继承树根节点

基础语法 —— json 库
- 用于处理 json 格式的字符串
- 字段后面的内容被称为 Tag，即标签，运行期间可以反射拿到
- json库依据 json Tag 的内容来完成json数据到结构体的映射
- 典型的声明式API设计
*/

func signUp(w http.ResponseWriter, r *http.Request) {
	req := &request.SignUpReq{}
	// ======这一块代码，但凡你要读json输入，就得来一遍【start】======
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		return
	}
	// ======这一块代码，但凡你要读json输入，就得来一遍【end】======

	// 返回一个虚拟的user id表示注册成功
	fmt.Fprintf(w, "%v", req)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is home")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is order")
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is createOrder")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is user")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is createUser")
}
