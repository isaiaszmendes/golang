package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // false
	x = true
	fmt.Println(x)   // valor atribuído -> true
	x = (10 < 100)   // true
	y := (10 == 100) // false
	z := (10 > 100)  // false
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
