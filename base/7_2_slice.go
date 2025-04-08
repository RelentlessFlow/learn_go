package main

import "fmt"

func main() {
	// slice 定义
	// 方式一
	var courses []string
	fmt.Printf("%T\r\n", courses)

	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")

	fmt.Println(courses[1])

	allCourses := []string{"go", "grpc", "http", "mysql", "redis"}
	courseSlice := allCourses[0:len(allCourses)]

	fmt.Println(courseSlice)

	// 方式二
	allCourses2 := make([]string, 3)
	allCourses2[0] = "c"
	allCourses2[1] = "c"
	allCourses2[2] = "c"
	//allCourses2[3] = "c" index out of range [3] with length 3

	fmt.Println(allCourses2)

	// 方式三
	var allCourses3 []string
	allCourses3 = append(allCourses3, "c")
	fmt.Println(allCourses3)

	// 切片访问
	fmt.Println(courseSlice[1])   // grpc
	fmt.Println(courseSlice[1:4]) // [grpc http mysql]
	fmt.Println(courseSlice[1:])  // [grpc http mysql redis]
	fmt.Println(courseSlice[:])   // [go grpc http mysql redis]

	courseSlice2 := []string{"go2", "grpc2", "http2", "mysql2", "redis2"}
	courseSlice3 := append(courseSlice, courseSlice2[2:]...)

	fmt.Println(courseSlice3)

	// 删除slice中的元素
	courseSlice2 = append(courseSlice2[:2], courseSlice2[3:]...)
	fmt.Println(courseSlice2)

	// 复制slice
	var courseSliceCopy = make([]string, len(courseSlice))
	copy(courseSliceCopy, courseSlice)
	fmt.Println(courseSliceCopy)

	// slice 是 值传递 还是 引用传递，是值传递
	var mySlice = []string{"1", "2", "3", "4", "5"}
	fmt.Println(len(mySlice), cap(mySlice)) // 5 5
	var mySlice2 = mySlice[1:2]
	fmt.Println(len(mySlice2), cap(mySlice2)) // 1 4
	mySlice2[0] = "100"
	mySlice2 = append(mySlice2, "6", "7", "8", "9")
	fmt.Println(len(mySlice2), cap(mySlice2)) // 5 8
	mySlice2[1] = "101"
	fmt.Println(mySlice)  // [1 100 3] 扩容之前，指针指向原数组
	fmt.Println(mySlice2) // [100 101 5 6]，扩容以后，指针指向新数组
}
