package main

import "fmt"

// Print the number 8 in decimal, binary and hexadecimal

func main() {
	num := 8
	// decimal
	fmt.Printf("%d\n", num) // 8
	// binary
	fmt.Printf("%b\n", num) // 1000
	// hexadecimal
	fmt.Printf("%#x\n", num) // 0x8
	// hexadecimal withouth the prefix
	fmt.Printf("%x\n", num) // 8,
	// contudo, é recomendado usar %#x para não confundir com decimal
}
