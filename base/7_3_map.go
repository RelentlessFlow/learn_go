package main

import "fmt"

func main() {
	// map 是一个key（索引）和 value（值）的无序集合，查询方便o(1)
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc工程师",
		"gin":  "gin深入理解",
	}

	// 取值
	fmt.Println(courseMap)
	fmt.Println(courseMap["grpc"])

	cv, cExist := courseMap["grpc"]
	fmt.Println(cv, cExist) // grpc工程师 true

	if _, cExist := courseMap["grpc"]; !cExist {
		fmt.Println("not in")
	} else {
		fmt.Println("in...")
	}

	// 放值
	courseMap["mysql"] = "mysql原理"
	fmt.Println(courseMap)

	// mao 若想使用，必须初始化，否则默认是nil
	var courseMap2 = map[string]string{}
	fmt.Println(courseMap2)

	// make是内置函数，主要用于初始化 slice、map、channel
	var courseMap3 = make(map[string]string, 3)
	fmt.Println(courseMap3)

	for k, v := range courseMap {
		// go go工程师 grpc grpc工程师 gin gin深入理解 mysql mysql原理
		fmt.Println(k, v)
	}

	for v := range courseMap {
		fmt.Println(v) // gin mysql go grpc
	}

	// 删除一个元素
	delete(courseMap, "mysql")
	delete(courseMap, "rpc")
	delete(courseMap, "rpc2") // 不会报错

	// 提示 map 不是线程安全的
}
