package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int
	var i8 int8 = 12
	var ui8 uint8 = uint8(i8)
	// float
	var f32 float32 = 3.14
	var i32 = int32(f32) // 3
	// str
	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println(err) // strconv.Atoi: parsing "12a": invalid syntax
	}
	var myi = strconv.Itoa(myint)
	fmt.Println(ui8, i32, myint, myi)

	f64, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println(err) // strconv.ParseFloat: parsing "3.1415a": invalid syntax
	}

	fmt.Println(f64) // 3.1415

	i64, err := strconv.ParseInt("-42", 8, 64)
	fmt.Println(i64) // -34

	// bool  0 true 转换后是true，其余是false
	parseBool, err := strconv.ParseBool("0")
	if err != nil {
		return
	}
	fmt.Println(parseBool) // false
	parseBool2, err := strconv.ParseBool("1")
	fmt.Println(parseBool2) // true
	parseBool3, err := strconv.ParseBool("3")
	fmt.Println(parseBool3) // false
	parseBool4, err := strconv.ParseBool("true")
	fmt.Println(parseBool4) // true
}
