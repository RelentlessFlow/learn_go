package main

import (
	"container/list"
	"fmt"
)

func main() {
	// list 链表
	myList := list.New()
	// 插入
	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("mysql")

	// 指定位置插入
	i := myList.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "grpc" {
			break
		}
	}
	myList.InsertBefore("gin", i)

	// 遍历
	for i := myList.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	// 删除指定元素
	myList.Remove(i)
}
