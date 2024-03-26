package main

import "fmt"

/**
基础语法——包声明
- 语法形式：packagexxxx
- 字母和下划线的组合
- 可以和文件夹不同名字
- 同一个文件夹下的声明一致

基础语法——包声明
- 引入包语法形式：import[alias]xxx
- 如果一个包引入了但是没有使用，会报错
- 匿名引入：前面多一个下划线

基础语法——main函数要点：
- 无参数、无返回值
- main方法必须要在main包里面
- `gorunmain.go`就可以执行
- 如果文件不叫`main.go`，则需要`gobuild`之后再`gorun`
*/

/**
- var，语法：varnametype=value
	- 局部变量
	- 包变量
	- 块声明
- 驼峰命名
- 首字符是否大写控制了访问性：大写包外可访问；
- golang支持类型推断

1.string类型——和别的语言没啥区别
2.基础类型——不必死记硬背，GolandIDE会提示你
3.切片——make,[i],len,cap,append
4.数组——和别的语言没啥区别
5.for,if,switch——和别的语言区别不大，IDE会提示你
*/
// Global首字母大写，全局可以访问
var Global = "全局变量"

// 首字母小写，只能在这个包里面使用
// 其子包也不能用
var local = "包变量"

// 块声明，一起声明多个变量
var (
	First  string = "abc"
	second int32  = 26
)

var aa = "hello"

//var aa = 1

const internal = "包内可访问"
const External = "包外可访问"

func main() {
	// fmt.Print("hello go")
	// print是builtin[执行内建的函数；内键指令；安装在内部的；装入的，内装式]包下的，可以直接使用，无需导包
	// print("hello go!")
	/**
	基础语法——string声明
	string：
	双引号引起来，则内部双引号需要使用\转义
	`反引号引起来，则内部`不需要\转义
	*/
	//print("He said:\" Hello Go!\"")
	//print(`He said: "Hello Go!"
	//go....`)

	/**
		基础语法——string长度
	- string的长度很特殊：
	  - 字节长度：和编码无关，用len(str)获取
	  - 字符数量：和编码有关，用编码库来计算
	Tip：如果你觉得字符串里边会出现非ASCII的字符，就记得用utf8库来计算“长度”
	*/
	// rune：如尼字母，理解成字符就行了
	//println(len("你好")) // 6
	//println(len("你好ab")) // 8
	//println(utf8.RuneCountInString("你好")) // 2
	//println(utf8.RuneCountInString("你好ab")) // 4

	/**
	基础语法——strings包
	- string的拼接直接使用+号就可以。注意的是，某些语言支持string和别的类型拼接，但是golang不可以
	- strings主要方法（你所需要的全部都可以找到）：
	  - 查找和替换
	  - 大小写转换
	  - 子字符串相关
	  - 相等
	*/
	//res := strings.Compare("a", "a")
	//println(res)

	/**
	基础语法——rune类型
	rune，直观理解，就是字符
	rune不是byte!
	rune本质是int32，一个rune四个字节
	rune在很多语言里面是没有的，与之对应的
	是，golang没有char类型。rune不是数字，
	也不是char，也不是byte！
	实际中不太常用
	type rune = int32
	*/

	/**
	基础语法——bool,int,uint,float家族
	bool:true,false
	int8,int16,int32,int64,int
	uint8,uint16,uint32,uint64,uint
	float32,float64
	*/
	//var b bool = false
	//println(b)
	//var i int = 6
	//println(i)
	//var ui uint = 6
	//println(ui)
	//var f float32 = 6.6
	//println(f)

	/**
	基础语法——byte类型
	byte，字节，本质是uint8
	对应的操作包在bytes上
	*/
	//var b byte = 255
	//println(b)

	/**
	基础语法——类型总结
	golang的数字类型明确标注了长度、有无符号
	golang不会帮你做类型转换，类型不同无法通过编译。
	也因此，string只能和string拼接
	golang有一个很特殊的rune类型，接近一般语言的char或者character的概念，非面试情况下，可以理解为rune=字符”
	string遇事不决找strings包
	*/
	//var name string = "hello"
	//println(name)
	//var char rune = 'h'
	//println(char) // 104

	/**
	基础语法——变量声明var
	var，语法：varnametype=value
	局部变量
	包变量
	块声明
	驼峰命名
	首字符是否大写控制了访问性：大写包外可访问；
	golang支持类型推断
	*/

	/**
	基础语法——变量声明:=
	- 只能用于局部变量，即方法内部
	golang使用类型推断来推断类型。数字会被理
	解为int或者float64。（所以要其它类型的数
	字，就得用var来声明）
	*/

	// int是灰色的，是因为golang自己可以做类型推断，它觉得你可以省略
	//var a int = 26
	//println(a)
	//
	//// 省略类型
	//var b = 26
	//println(b)
	//
	//var c uint = 26
	//print(c)

	// 无法通过编译，因为golang是强类型语言，并且不会帮你做任何的转换
	// println(a == c)

	//a := 26
	//println(a)
	//b := "hello"
	//println(b)
	//
	//f := 0.26
	//println(f)

	/**
	基础语法——变量声明易错点
	- 变量声明了没有使用
	- 类型不匹配
	- 同作用域下，变量只能声明一次
	*/
	//aa := 1 // 虽然包外面已经有一个aa了，但是这里从包变成了局部变量
	//println(aa)
	//
	//var bb = 2
	////var bb = "two" // 重复声明，也会导致编译不通过
	//
	//bb = 6 // ok，没有重复声明，只是赋值了新的值
	//// bb := 6 // 不行，因为:=就是声明并且赋值的简写，相当于重复声明了bb
	//println(bb)

	/**
	基础语法——常量声明const
	首字符是否大写控制了访问性：大写包外可访问；
	驼峰命名
	支持类型推断
	无法修改值
	*/
	//const a = "hello"
	//println(a)

	//var greeting = Fun0(`luffy`)
	//println(greeting)

	//var age,name = Fun2("19","luffy")
	//println(age)
	//println(name+" 空字符串")

	/**
	Golang语法——方法调用
	使用_忽略返回值
	*/
	//var age, _ = Fun2("19", "luffy")
	//println(age)

	/**
	Golang语法——方法声明与调用总结
	•golang支持多返回值，这是一个很大的不同点
	•golang方法的作用域和变量作用域一样，通过大小写控制
	•golang的返回值是可以有名字的，可以通过给予名字让调用方清楚知道你返回的是什么
	*/

	//name := "luffy"
	//age := 20
	//// Sprintf格式化字符串并返回
	//introduction := fmt.Sprintf("hello I am %s, %d years old", name, age)
	//println(introduction)
	//
	//// 直接输出字符串
	// fmt.Printf("hello I am %s", name)
	// formativeOutput()
	// arrOpt()
	// sliceOpt()
	// subSliceOpt()
	// shareSlice()
	// forOpt1()
	// forOpt2()
	// forOpt3()
	// IfUsingNewVariable(100,200)
	// ChooseFruit("apple")
	// fake := FakeFish{}
	// fake无法调用原来Fish的方法
	// 这一句会编译报错
	// fake.Swim()
	// fake.FakeSwim()

	// 转换为Fish
	// td := Fish(fake)
	// 真的变成了鱼
	// td.Swim()

	// sFake := StrongFakeFish{}
	// 这里就是调用了自己的方法
	// sFake.Swim()

	// td = Fish(sFake)
	// 真的变成了鱼
	// td.Swim()

	//a := A{}
	// 当结构体内成员相同时，可以进行类型转换
	//b := B(a)
	//var n News = fakeNews{
	//	Name: "hello",
	//}
	//n.Report()
}

type News struct {
	Name string
}

func (d News) Report() {
	fmt.Println("I am news: " + d.Name)
}

type fakeNews = News

/**
基础语法——typeA=B
基本语法:type TypeA = TypeB
别名，除了换了一个名字，没有任何区别
*/

type A struct {
	field string
}

type B struct {
	field string
}

/*
*
基础语法——type A B
基本语法:type TypeA TypeB
我一般是，在我使用第三方库又没有办法修改源码的情况下，
又想在扩展这个库的结构体的方法，就会用这个
Tip：这个不用记，属于那种看上去很复杂，但是实际你根本不会这么写的东西。
*/
type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼，假装自己是鸭子 \n")
}

// 定义了一个新类型，注意是新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎嘎 \n")
}

type StrongFakeFish Fish

func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼，嘎嘎嘎 \n")
}

/*
*
基础语法——switch
switch和别的语言差不多
switch后面可以是基础类型和字符串，或者满足特定条件的结构体最大的差别：
终于不用加break了！
Tip：大多数时候，switch后面只会用基础类型或者字符串
*/
func ChooseFruit(fruit string) {
	switch fruit {
	case "apple":
		fmt.Println("this is an apple")
	case "banana", "orange":
		fmt.Println("this is a banana or an orange")
	default:
		fmt.Println("this is an unknown fruit:{}", fruit)
	}
}

/**
基础语法——if-else
带局部变量声明的if-else：
1.distance只能在if块，或者后边所有的else块里面使用
2.脱离了if-else块，则不能再使用
*/

func IfUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("距离太远，不来了：%d \n", distance)
	} else {
		fmt.Printf("距离不远，来一趟：%d \n", distance)
	}

	// 这里不能访问 distance
	// fmt.Printf("距离是：%d \n",distance)
}

/**
基础语法——for
for和别的语言差不多，有三种形式：1.for{}，类似while的无限循环
2.fori，一般的按照下标循环
3.forrange最为特殊的range遍历
4.break和continue和别的语言一样
*/

func forOpt3() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range arr {
		fmt.Printf("%d => %d \n", index, value)
	}
}

func forOpt2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func forOpt1() {
	arr := []int{7, 8, 9}
	index := 0
	for {
		if index == 1 {
			break
		}
		fmt.Printf("%d => %d \n", index, arr[index])
		index++
	}
	fmt.Println("for loop end")
}

/*
*
基础语法——共享底层（optional）
核心：共享数组
子切片和切片究竟会不会互相影响，就抓住一点：它们是不是还共享数组？
什么意思？就是如果它们结构没有变化，那肯定是共享的；
但是结构变化了，就可能不是共享了
有余力的同学可以运行一下ShareSlice()
*/
func shareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	// s1: [1 2 3 4],len: 4,cap: 4
	fmt.Printf("s1: %v,len: %d,cap: %d \n", s1, len(s1), cap(s1))
	// s2: [3 4],len: 2,cap: 2
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	// s1: [1 2 99 4],len: 4,cap: 4
	fmt.Printf("s1: %v,len: %d,cap: %d \n", s1, len(s1), cap(s1))
	// s2: [99 4],len: 2,cap: 2
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	// s1: [1 2 99 4],len: 4,cap: 4
	fmt.Printf("s1: %v,len: %d,cap: %d \n", s1, len(s1), cap(s1))
	// s2: [99 4 199],len: 3,cap: 4
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))

	s2[1] = 1999
	// s1: [1 2 99 4],len: 4,cap: 4
	fmt.Printf("s1: %v,len: %d,cap: %d \n", s1, len(s1), cap(s1))
	// s2: [99 1999 199],len: 3,cap: 4
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))
}

/**
基础语法——子切片
数组和切片都可以通过[start:end]的形式来获取子切片：
1.arr[start:end]，获得[start,end)之间的元素
2.arr[:end]，获得[0,end)之间的元素
3.arr[start:]，获得[start,len(arr))之间的元素
包左不包右

如何理解切片
最直观的对比：ArrayList，即基于数组的List的实现，切片的底层也是数组
跟ArrayList的区别：
1.切片操作是有限的，不支持随机增删（即没有add,delete方法，需要自己写代码）
2.只有append操作
3.切片支持子切片操作，和原本切片是共享底层数组
***切片支持子切片操作，和原本切片是共享底层数组的***
*/

func subSliceOpt() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	// s2: [2 3],len: 2,cap: 5
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	// s3: [3 4 5 6],len: 4,cap: 4
	fmt.Printf("s3: %v,len: %d,cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	// s4: [1 2 3],len: 3,cap: 6
	fmt.Printf("s4: %v,len: %d,cap: %d \n", s4, len(s4), cap(s4))
}

/**
            数组       切片
直接初始化    支持       支持
make        不支持      支持
访问元素     arr[i]     arr[i]
len        长度         已有元素个数
cap        长度        容量
append     不支持      支持
是否可以扩容 不可以      可以

Tip：遇事不决用切片，基本不会出错
*/

/*
*
切片,语法：[]type
1.直接初始化
2.make初始化:make([]type,length,capacity)
3.arr[i]的形式访问元素
4.append追加元素
5.len获取元素数量
6.cap获取切片容容量
7.推荐写法：s1:=make([]type,0,capacity)

Tip：初学的时候不必关心什么时候扩容，什么时候不扩容
超过capacity并且数量小于1024，翻倍扩容，大于1024每次扩容25%
*/
func sliceOpt() {
	// 直接初始化了4个元素的切片
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v,len: %d,cap: %d \n", s1, len(s1), cap(s1))

	// 创建了一个包含三个元素，容量为4的切片
	s2 := make([]int, 3, 4)
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))
	// fmt.Printf("s2[4]: %d",s2[4])

	// 后边添加一个元素，没有超出容量限制，不会发生扩容
	s2 = append(s2, 7)
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))

	// 触发扩容
	s2 = append(s2, 8)
	fmt.Printf("s2: %v,len: %d,cap: %d \n", s2, len(s2), cap(s2))

	// 只传一个参数，表示创建一个含有4个元素，容量也为4个元素的切片
	s3 := make([]int, 4)
	fmt.Printf("s3: %v,len: %d,cap: %d \n", s3, len(s3), cap(s3))

	// 按下标索引
	fmt.Printf("s3[2]：%d", s3[2])

	// 超出下标范围，直接奔溃
	// runtime error: index out of range [99] with length 4
	// fmt.Printf("s3[99]：%d",s3[99])

}

/**
基础语法——数组和切片
数组和别的语言的数组差不多，语法是：[cap]type，[容量]类型
1.初始化要指定长度（或者叫做容量）
2.直接初始化
3.arr[i]的形式访问元素
4.len和cap操作用于获取数组长度【Tip：数组的len和cap结果是一样的，就是数组的长度】
*/

func arrOpt() {
	// 直接初始化一个三个元素的数组，大括号里面多一个少一个都编译不通过
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v,len: %d,cap: %d", a1, len(a1), cap(a1))
	fmt.Println()
	// 初始化一个三个元素的数组，所有元素都是0值
	var a2 [3]int
	fmt.Printf("a2: %v,len: %d,cap: %d", a2, len(a2), cap(a2))
	fmt.Println()
	// a1 = append(a1,12) 数组不支持append操作

	// 按下标索引
	fmt.Printf("a1[1]:%d", a1[1])

	// 超出下标范围，直接崩溃，编译不通过
	// fmt.Printf("a1[99]:d%",a1[99])
}

/**
fmt格式化输出
fmt包有完整的说明
掌握常用的：%s,%d,%v,%+v,%#v
不仅仅是`fmt`的调用，所有格式化字符串的API都可以用
因为golang字符串拼接只能在string之间，
所以这个包非常常用
*/

type User struct {
	Name string
	Age  int
}

func formativeOutput() {
	user := &User{
		Name: "Luffy",
		Age:  20,
	}
	fmt.Printf("v => %v \n", user)
	fmt.Printf("+v => %+v \n", user)
	fmt.Printf("#v => %#v \n", user)
	fmt.Printf("T => %T \n", user)
}

/**
Golang语法——方法声明
四个部分：
•关键字func
方法名字：首字母是否大写决定了作用域
•参数列表：[<name type>]
•返回列表:[<name type>]

Golang语法——方法声明（推荐写法）
•参数列表含有参数名
•返回值不具有返回值名
*/

// Fun0只有一个返回值，不需要括号括起来
func Fun0(name string) string {
	return "Hello, " + name
}

// Fun1多个参数，多个返回值，参数有名字，但是返回值没有名字
func Fun1(a string, b int) (int, string) {
	return 0, "你好"
}

/*
*
Golang语法——方法声明（看看就好）
Fun2的返回值具有名字，可以在内部直接复制，然后返回
也可以忽略age,name.知己返回别的。
*/
func Fun2(s string, b string) (age int, name string) {
	age = 19
	// name = "luffy"
	return
	// return 19,"luffy" // 这样返回也可以
}

// Fun3多个参数具有相同类型放在一起，可以只写一次类型
func Fun3(a, b, c string, abc, bcd int, p string) (d, e int, g string) {
	d = 15
	e = 16
	g = "你好"
	return
}

/**
课后练习
•计算斐波那契数列
•实现切片的Add和Delete方法
去leetcode上试试（先看答案，再尝试用go写出来）
•我们课上用了很多fmt来格式化字符串，那么如何输出
3.1保留两位小数的数字
3.2将[]byte输出为16进制
预习type的用法
*/
