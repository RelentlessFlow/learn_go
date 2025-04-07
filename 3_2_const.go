package main

func main() {
	// 显式定义
	const PI float32 = 3.1415926

	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	const (
		A = "AA"
		B = "BB"
		C
		D
		E = 'E'
		F
	)

	// 若常量未赋值，沿用临值
	println(A, B, C, D, E)
}
