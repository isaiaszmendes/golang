package main

import "fmt"

var x bool
var num int

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
	fmt.Printf("=> %T  and %v\n", num, num) // bool

	space32 := []byte(" ")
	space160 := rune(160)
	fmt.Printf("=> %v\n", space32)
	fmt.Printf("=> Charcode:%c.\n", space160)
}
