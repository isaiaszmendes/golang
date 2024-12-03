package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x) // int
	fmt.Printf("%T\n", y) // string
	fmt.Printf("%T\n", z) // bool
}
