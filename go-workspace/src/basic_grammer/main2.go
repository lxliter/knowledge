package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
*
1.http库
2.基础语法type
3.Server与Context抽象
4.简单支持RESTFulAPI
*/
func main() {
	/**
	func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		DefaultServeMux.HandleFunc(pattern, handler)
	}
	*/
	http.HandleFunc("/", readBodyOnce)
	http.HandleFunc("/getBodyIsNil", getBodyIsNil)
	http.HandleFunc("/queryParams", queryParams)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/header", header)
	http.HandleFunc("/form", form)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("web server start failure %v", err)
	}
}

/**
要点总结：http库使用
Body和GetBody：重点在于Body是一次性的，而GetBody默认情况下是没有，一般中间件会考虑帮你注入这个方法
URL：注意URL里面的字段的含义可能并不如你期望的那样
Form：记得调用前先用ParseForm，别忘了请求里面加上http头
 */

/*
*
http库——Form
Form和ParseForm
要先调用ParseForm
建议加上Content-Type:application/x-www-form-urlencoded

before parse form map[]
after parse form map[age:[20] name:[luffy]]
*/
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v \n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "after parse form error %v \n", r.Form)
	}
	fmt.Fprintf(w, "after parse form %v \n", r.Form)
}

/*
*
http库——RequestHeader
header大体上是两类，一类是http预定义的；一类是自己定义的
Go会自动将header名字转为标准名字——其实就是大小写调整
***一般用X开头来表明是自己定义的***，比如说X-mycompany-your=header
header is map[
Accept:[*\*]
Accept-Encoding:[gzip, deflate, br]
Connection:[keep-alive]
Content-Length:[23]
Content-Type:[text/plain]
Postman-Token:[14bf32cb-cf66-437d-a8c2-5ad6c1763b31]
User-Agent:[PostmanRuntime/7.28.3]
X-Luffy-Myapp:[test-app]]
*/
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v \n", r.Header)
}

/*
*
http库——RequestURL
包含路径方面的所有信息和一些很有用的操作

	type URL struct {
		Scheme 		string
		Opaque 		string
		User   		*Userinfo
	    Host   		string
	    Path   		string
	    RawPath 	string
	    ForceQuery  bool
	    RawQuery    string
	    Fragment    string
	    RawFragment string
	}

http库——RequestURL
URL里面Host不一定有值
r.Host一般都有值，是Host这个header的值
RawPath也是不一定有值
Path肯定有
Tip：实际中记得自己输出来看一下，确认有没有
localhost:8080/wholeUrl?a=123&a=456&b=789

	{
	    "Scheme": "",
	    "Opaque": "",
	    "User": null,
	    "Host": "",
	    "Path": "/wholeUrl",
	    "RawPath": "",
	    "OmitHost": false,
	    "ForceQuery": false,
	    "RawQuery": "a=123&a=456&b=789",
	    "Fragment": "",
	    "RawFragment": ""
	}
*/
func wholeUrl(w http.ResponseWriter, r *http.Request) {
	// func Marshal(v any) ([]byte, error)
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

/*
*
http库——RequestQuery
除了Body，我们还可能传递参数的地方是Query
也就是http://xxx.com/your/path?id=123&b=456
所有的值都被解释为字符串，所以需要自己解析为数字等
localhost:8080/queryParams?a=123&a=456&b=789
*/
func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	// query is map[a:[123 456] b:[789]]
	fmt.Fprintf(w, "query is %v \n", values)
}

/*
*
http库——RequestBody-GetBody
Body：只能读取一次，意味着你读了别人就不能读了；别人读了你就不能读了；
GetBody：原则上是可以多次读取，但是在原生的http.Request里面，这个是nil
在读取到body之后，我们就可以用于反序列化，比如说将json格式的字符串转化为一个对象等
*/
func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil \n")
	} else {
		fmt.Fprint(w, "GetBody is not nil")
	}
}

/*
*
http库——RequestBody
Body：只能读取一次，意味着你读了
别人就不能读了；别人读了你就不能读了；
*/
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	// 类型转换，将[]byte转换为string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the dataa one more time error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

/**
http库——Request概览
•Body和GetBody
•URL
•Method
•Header
•Form
*/

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是首页")
}
