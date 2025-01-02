package main

import "fmt"

func main() {
	a := 100

	fmt.Printf("%d \t %b \t %#x\n", a, a, a) // 100 1100100 0x64
	// 0x (é um prefixo que indica que o número está em hexadecimal)
}

// Base - 10: decimal, 0-9 (Foramatção %d)
// Base - 2: binário, 0-1 (Formatação %b)
// Base - 16: hexadecimal, 0-F (Formatação %#x)

// O sistema hexadecimal é uma forma interante para o computador
// pois é mais fácil de converter para binário
