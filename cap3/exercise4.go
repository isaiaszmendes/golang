package main

import "fmt"

type idade int

var x idade

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)
}
