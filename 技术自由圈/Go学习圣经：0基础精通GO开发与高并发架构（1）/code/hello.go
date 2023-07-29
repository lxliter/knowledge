package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")

    var input string
    fmt.Scanln(&input)

    fmt.Println("你的输入是：%s", input)

}