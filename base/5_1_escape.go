package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 字符串长度计算，如果只有英文，可以直接len，否则用[]rune 转换一下
	name := "imooc体系课学习"
	bytes := []rune(name)

	fmt.Println(len(bytes)) // 10
	fmt.Println(len(name))  // 20

	// 转义符
	escapeStr := "go\"体系课\""
	fmt.Println(escapeStr)

	templateStr := `go"体系课"`
	fmt.Println(templateStr)

	// 格式化输出
	var username []string = []string{"bill", "gates"}
	age := 18
	address := "北京"
	mobile := "188888888"

	fmt.Println("用户名：" + strings.Join(username, " ") + "，年龄：" + strconv.Itoa(age) + "，地址：" + address + "，电话：" + mobile)
	fmt.Printf("用户名：%s，年龄：%d，地址：%s，电话：%s\r\n", username, age, address, mobile)

	userMsg := fmt.Sprintf("用户名：%s，年龄：%d，地址：%s，电话：%s\r\n", username, age, address, mobile)

	// %s:		[bill gates]
	// %#v:		[]string{"bill", "gates"}
	// %T		[]string

	fmt.Println(userMsg)
}
