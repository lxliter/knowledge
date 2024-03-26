package main

import (
	"fmt"
)

func main() {
	//e1 := &Example{}
	///**
	//指针：结构体初始化方式1：&{ 0}
	//*/
	//fmt.Printf("结构体初始化方式1：%v \n", e1)
	//
	//e2 := Example{}
	///**
	//结构体初始化方式2：{ 0}
	//*/
	//fmt.Printf("结构体初始化方式2：%v \n", e2)
	//
	//e3 := new(Example)
	///**
	//指针：结构体初始化方式3：&{ 0}
	//*/
	//fmt.Printf("结构体初始化方式3：%v \n", e3)
	//e1.m1()
	//e1.m2()

	/**
	当你声明成这样的时候，go就帮你分配好内存了
	不用担心空指针的问题，因为它压根就不是指针
	*/
	//var e1 Example
	//e1.m1()
	//
	//// e2就是一个指针
	//var e2 *Example
	//// 这边会直接panic掉
	//// panic: runtime error: invalid memory address or nil pointer dereference[废弃， 解除引用]
	//e2.m1()

	//s1 := Student{
	//	name: "luffy",
	//	age:  20,
	//}
	//
	//s1.sayHi()
	//
	//// 初始化按字段顺序赋值，不建议使用
	//s2 := Student{"lucy", 19}
	//s2.sayHi()
	//
	//// 先初始化，再单独赋值
	//s3 := Student{}
	//s3.name = "lily"
	//s3.age = 18
	//s3.sayHi()
	// handlePointer()

	//s1 := Struct1{}
	//s1.v1 = "oldv1"
	////s2 := Struct2{
	////	v2: "oldv2",
	////}
	//s1.s2.v2 = "oldv2"
	//fmt.Printf("s1 old value %v \n", s1)
	//// changeValue1(s1)
	//changeValue2(s1)
	//fmt.Printf("s1 new value %v \n", s1)

	// 因为stu是结构体，所以方法调用的时候，它的数据是不会变的
	//stu := Student{
	//	name: "luffy",
	//	age:  18,
	//}
	//stu.ChangeName("luffy changed")
	//// stu changeName {luffy 18}
	//fmt.Printf("stu changeName %v \n", stu)
	//
	//stu.ChangeAge(20)
	//// stu changeAge {luffy 20}
	//fmt.Printf("stu changeAge %v \n", stu)
	//
	//stu1 := &Student{
	//	name: "luffy",
	//	age:  18,
	//}
	//stu1.ChangeName("luffy changed")
	//// stu1 changeName &{luffy 18}
	//fmt.Printf("stu1 changeName %v \n", stu1)
	//
	//stu1.ChangeAge(21)
	//// stu1 changeAge &{luffy 21}
	//fmt.Printf("stu1 changeAge %v \n", stu1)
}

/**
基础语法——结构体如何实现接口？
当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。
当一个结构体具备接口的所有的方法的时候，它就实现了这个接口
 */

/**
基础语法——方法接收器用哪个？
- 设计不可变对象，用结构体接收器
- 其它用指针
总结：遇事不决用指针
 */

/**
基础语法——方法接收器
Tip：***结构体和指针之间的方法可以互相调用***
 */


/**
基础语法——方法接收器
结构体接收器内部永远不要修改字段
*/
// 结构体接收器
func (s Student) ChangeName(newName string) {
	s.name = newName
}

// 指针接收器
func (s *Student) ChangeAge(newAge int) {
	s.age = newAge
}

func changeValue2(s Struct1) {
	s.v1 = "newv1"
	s.s2.v2 = "newv2"
}

func changeValue1(s *Struct1) {
	s.v1 = "newv1"
	s.s2.v2 = "newv2"
}

type Struct1 struct {
	v1 string
	s2 Struct2
}

type Struct2 struct {
	v2 string
}

/*
*
基础语法——结构体自引用
结构体内部引用自己，只能使用指针
准确来说，在整个引用链上，如果构成循环，那就只能用指针
*/
type Node struct {
	// 自引用只能使用指针
	// Invalid recursive type 'Node
	//left Node
	//right Node

	left  *Node
	right *Node

	// 这个也会报错
	// nn NodeNode
}

type NodeNode struct {
	node Node
}

/*
*
基础语法——指针
和C，C++一样，*表示指针，&取地址
如果声明了一个指针，但是没有赋值，那么它是nil
*/
func handlePointer() {
	// 指针用*表示
	var sPointer *Student = &Student{}
	sPointer.name = "luffy"
	sPointer.age = 20
	// 解引用，得到结构体
	var sStruct Student = *sPointer
	sStruct.sayHi()

	// 只是声明了，但是没有使用
	var stu *Student
	if stu == nil {
		fmt.Println("stu is nil")
	}
}

/*
*
基础语法——字段赋值
*/
type Student struct {
	name string
	age  int
}

func (s Student) sayHi() {
	fmt.Printf("I am %s, %d years old \n", s.name, s.age)
}

/*
*
Go没有构造函数！！
初始化语法：Struct{}
获取指针：&Struct{}
获取指针2：new(Struct)：new可以理解为Go会为你的变量分配内存，并且把内存都置为0
*/
type Example struct {
	f1 string
	f2 int
}

func (e Example) m1() {
	fmt.Printf("invoke m1 \n")
}

func (e Example) m2() {
	fmt.Printf("invoke m2 \n")
}
