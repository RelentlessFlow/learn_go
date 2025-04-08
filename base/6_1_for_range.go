package main

import "fmt"

func main() {
	name := "Hello World"
	for index, value := range name {
		fmt.Printf("%d %c \n", index, value) // 0 h ...
	}

	for index := range name {
		fmt.Printf("%d \n", index) // 0 1 2 3 4 ...
	}

	/**
	字符串	字符串的索引key 字符串对应的索引的拷贝				不写key，返回的是索引
	数组 	数组的索引 	  索引对应的值的拷贝   				不写key，返回的是索引
	切片 	切片的索引 	  索引对应的值的拷贝   				不写key，返回的是索引
	map 	map的key 	  value 返回的是 key 对应的值的拷贝   	不写key，返回的是map的值
	channel				  value 返回的是 channel 接受的值
	*/

	zhName := "你好World"
	zhNameRune := []rune(zhName) // 直接用 index打印中文会失效，需要转为rune

	for i := 0; i < len(zhNameRune); i++ {
		fmt.Printf("%c", zhNameRune[i])
	}

	fmt.Println()

	for _, value := range zhName {
		fmt.Printf("%c", value)
	}
}
