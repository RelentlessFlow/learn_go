package main

import "fmt"

func main() {
	// 集合类型 数组、切片、map、list
	var courses1 [3]string
	var courses2 [4]string
	var tStr1 = fmt.Sprintf("%T", courses1)
	var tStr2 = fmt.Sprintf("%T", courses2)

	fmt.Println(tStr1, tStr2) // [3]string [4]string

	courses1[0] = "go"
	courses1[1] = "grpc"
	courses1[2] = "gin"

	courses3 := [3]string{"go", "grpc", "gin"}
	courses4 := [3]string{2: "gin"}
	courses5 := [...]string{"go", "grpc", "gin"}
	courses6 := [...]string{}
	courses7 := []string{}
 
	fmt.Println(courses3, courses4) // [go grpc gin] [  gin]
	fmt.Printf("%T\n", courses5)    // [3]string
	fmt.Printf("%T\n", courses6)    // [0]string
	fmt.Printf("%T\n", courses7)    // []string 这玩意是个切片

}
