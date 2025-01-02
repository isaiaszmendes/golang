package main

import "fmt"

// Crie um programa que:
// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal

func main() {
	var x int = 10
	fmt.Printf("%d\n", x)  // 10
	fmt.Printf("%b\n", x)  // 1010
	fmt.Printf("%#x\n", x) // 0xa
	y := x << 1
	fmt.Printf("%d\n", y)  // 20
	fmt.Printf("%b\n", y)  // 10100
	fmt.Printf("%#x\n", y) // 0x14

}
