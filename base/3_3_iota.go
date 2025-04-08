package main

import "fmt"

func main() {
	const (
		ERR1 = iota + 1
		ERR2        // iota编译器会自增
		ERR3 = "ha" // 如果中端了iota，必须显式恢复
		ERR4
		ERR5 = iota
	)

	fmt.Println(ERR1, ERR2, ERR3, ERR5) // 1 2 ha 4

	const (
		ERR_NEW1 = iota
	) // iota 会在遇到const关键字重置为0

	fmt.Println(ERR_NEW1) // 0
}
