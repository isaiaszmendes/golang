package main

import "fmt"

func main() {
	x := 10
	y := x >> 1
	fmt.Printf("%b\n", x) // 1010
	fmt.Printf("%b\n", y) // 101
	x = 1
	y = x << 1
	fmt.Printf("%b\n", x) // 1
	fmt.Printf("%b\n", y) // 10
}
