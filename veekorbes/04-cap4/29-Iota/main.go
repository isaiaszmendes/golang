package main

import "fmt"

// Iota é uma constante que é incrementada automaticamente

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	fmt.Println(A) // 0
	fmt.Println(B) // 1
	fmt.Println(C) // 2
}
