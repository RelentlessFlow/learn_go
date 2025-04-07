package main

import "fmt"

func main() {
	// int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	var i int       // 动态类型
	i16 = int16(i8) // 类型转换
	// float
	var f32 float32
	var f64 float64
	// byte -> uint8 无符号数
	var c byte // 存放字符，asic码
	c = 'a'
	fmt.Println(c)      // 97
	fmt.Printf("%c", c) // a
	// rune 字符
	var r rune = '张'
	fmt.Println(r)      // a24352
	fmt.Printf("%c", r) // 张
	fmt.Println(
		i8, i16, i32, i64, i,
		u8, u16, u32, u64,
		f32, f64,
	)
	// string
	var str string = "Hello World"
	fmt.Println(str)

}
