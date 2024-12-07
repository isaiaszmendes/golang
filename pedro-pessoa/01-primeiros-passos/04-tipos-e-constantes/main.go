package main

import (
	"fmt"
	"strconv"
)

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8
// rune // alias for int32

// float32 float64

// complex64 complex128

func main() {
	var x uint8 = 100

	imprimeByte(x) // Deu certo porque uint8 é um alias para byte
	// Dois nomes diferentes porem que representam o mesmo tipo

	// conversão de tipos
	var xFloat float32 = float32(x)
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", xFloat, xFloat)
	xString := string(x)
	fmt.Printf("%v, %T\n", xString, xString) // d, string (d) é o valor do código ASCII 100
	// imprime o nome Isaias
	fmt.Println(string(73) + string(115) + string(97) + string(105) + string(97) + string(115))
	// imprime o nome Isaías com acento
	fmt.Println(string(73) + string(115) + string(97) + string(237) + string(97) + string(115))

	// conversão de int para string
	var y int = 100
	yString := strconv.FormatInt(int64(y), 10)
	fmt.Printf("%v, %T\n", yString, yString)

	const z = 30 // Chamamos isso de untyped constant
	fmt.Printf("%v, %T\n", z, z)
	takeInt32(z)
	takeInt64(z)
	const f = 3.14
	takeInt32(f) // Não compila porque f é float64 e takeInt32 espera um int32
}

func imprimeByte(b byte) {
	fmt.Println(b)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}
