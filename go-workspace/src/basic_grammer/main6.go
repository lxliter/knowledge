package main

import "fmt"

/**
组合可以是接口组合，也可以是结构体组合。
结构体也可以组合接口
*/

// swimming会游泳
type Swimming interface {
	Swim()
}

type Duck interface {
	// 鸭子是会游泳的，所以这里组合了它
	Swimming
}

type Base struct {
	Name string
}

func (b *Base) Swim() {
}

type Concrete1 struct {
	Base
}

type Concrete2 struct {
	*Base
}

func (b *Base) SayHello() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}

func (c Concrete1) SayHello() {
	// c.Name直接访问了Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)
	// 这样也是可以的
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)
	// 调用了被组合的
	c.Base.SayHello()
}

/**
Http Server —— 组合与重写
- Go 没有重写
- main 函数会输出 I am Parent
- 而在典型的支持重写的语言，如Java，我们可以期望它输出 I am Son
Tip：当你写下类似继承的代码的时候，千万要先试试它会调过去哪个方法
*/

type Parent struct {
}

func (p Parent) Name() string {
	return "Parent"
}

func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name())
}

type Son struct {
	Parent
}

func (s Son) Name() string {
	return "Son"
}

func (s Son) SayHello() {
	fmt.Println("I am " + s.Name())
}

func main() {
	//c := &Concrete1{
	//	Base: Base{
	//		Name: "Luffy",
	//	},
	//}
	//c.SayHello()

	son := Son{
		Parent{},
	}
	son.SayHello()
}

