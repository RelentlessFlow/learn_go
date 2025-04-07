package main

import "fmt"

// 一、全局变量
var gv = "boddy"
var age = 18
var ok bool

// 多变量定义
var (
	gv1 = 10
	gv2 = false
)

func main() {
	// 二、局部变量
	// 方法 1
	var v1 int
	v1 = 1
	// 方法 2
	var v2 = 1
	// 方法 3（常用，不能用于全局变量定义）
	v3 := 1
	// 局部变量定义了不使用是不行的，会报错
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	// 多变量定义
	var user1, user2, user3 = "body1", 1, "body3"
	fmt.Println(user1, user2, user3)

	// 三、变量都有零值（默认值）
	var num1 int
	var str string
	var bo bool
	fmt.Println(num1, str, bo)
}
