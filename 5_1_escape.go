package main

import "fmt"

func main() {
	// 字符串长度计算，如果只有英文，可以直接len，否则用[]rune 转换一下
	name := "imooc体系课学习"
	bytes := []rune(name)

	fmt.Println(len(bytes)) // 10
	fmt.Println(len(name))  // 20

	// 转义符
	courseName := "go\"体系课\""
	fmt.Println(courseName)
}
