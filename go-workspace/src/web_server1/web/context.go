package web

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
*
Context抽象与实现
- 读取数据
- 写入响应
*/
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

func (c *Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

/*
*
Http Server —— 写入响应
这里有个小差异，是我们不再使用 fmt，而是直接使用 Write 方法
*/
func (c *Context) WriteJson(status int, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(bs)
	if err != nil {
		return err
	}
	c.W.WriteHeader(status)
	return nil
}

/*
*
Http Server —— 进一步封装
提供辅助方法
注意！它不是 Context本身必须要提供的方法！
即如果你在设计真实的web 框架的时候，你需要考虑清楚，究竟要不要提供这种辅助方法
Tip：严格来说，WriteJson也是辅助方法
http库里面提前定义好了各种响应码
*/
func (c *Context) OkJson(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}

func (c *Context) SystemErrJson(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, data)
}

func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest, data)
}
