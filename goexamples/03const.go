package main

const (
	first  = iota + 6
	second = 2 << iota //bit shift 2 by 1
)
const (
	zero = iota
	one
	test
)

func main() {
	println(first, second)
	println(zero, one, test)
}