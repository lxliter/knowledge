package main

import "fmt"

func main() {
	// createMap()
	numMap := make(map[int]int)
	numMap[0] = 0
	numMap[1] = 1
	numMap[2] = 2
	forMap(numMap)
}

/*
*
基础语法 —— map 遍历
for key, val := range m {}
Go 一个 for 打天下
Go 的 map 的遍历，顺序是不定的
*/
func forMap(numMap map[int]int) {
	for key, val := range numMap {
		fmt.Printf("%d=>%d \n", key, val)
	}
}

/*
*
基础语法 —— map
- 基本形式：map[KeyType]ValueType
- 创建 make 命令，或者直接初始化
- 取值：val, ok := m[key]
- 设值：m[key]=val
- key 类型：“可比较”类型
Tip：编译器会告诉你能不能做 key
Tip：尽量用基本类型和string做key，不要和自己过不去
*/
func createMap() {
	// 创建一个预估容量为2的map
	//m := make(map[string]string, 2)
	//
	//// 没有指定预估容量
	//m1 := make(map[string]string)

	// 直接初始化
	//m2 := map[string]string{
	//	"Tom": "Jerry",
	//}

	// 赋值
	//m1["hello"] = "world"
	//// 取值
	//v1 := m1["hello"]
	//println(v1)
	//
	//val, ok := m["invalid_key"]
	//if !ok {
	//	println("key not found")
	//}
	//println(val)

	numMap := make(map[int]int)
	numMap[1] = 1
	numMap[2] = 2
	v1 := numMap[1]
	println(fmt.Sprintf("v1=%d", v1))
	v2 := numMap[2]
	println(fmt.Sprintf("v2=%d", v2))

	v3, ok := numMap[3]
	if !ok {
		println(fmt.Sprintf("key[%d] not found", 3))
	}
	println(fmt.Sprintf("v3=%d", v3))
}
