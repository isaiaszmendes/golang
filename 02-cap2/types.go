package main

import "fmt"

type hotdog int
var b hotdog
// "types is life"
func main() {
	x := 10
	fmt.Printf("%T \n", x) // int
	fmt.Printf("%T", b) // main.hotdog

	// Não posso fazer isso pq um é hotdog e o outro é int, mesmo ambos terem os mesmo tipo subjacente
	//! x = b

	//* Jeito certo
	// Conversion

	x = int(b)

	z := hotdog(x)
	

	fmt.Println("\n--------------")
	fmt.Printf("%T \n", x) 
	fmt.Printf("%T\n", b) 
	fmt.Printf("%T\n", z) //main.hotdog

}