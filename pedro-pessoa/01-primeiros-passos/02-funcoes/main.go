package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(somar(1, 2))
	fmt.Println(subtrair(1, 2))
	fmt.Println(swap(1, 2)) // Deve imprimir 2 1
	a, b := swap(3, 4)
	fmt.Println(a, b) // Deve imprimir 4 3

	divisao, resto := dividir(10, 3)
	fmt.Println(divisao, resto)       // Deve imprimir 3 1
	fmt.Println(sumHiguerOrder(2)(3)) // Deve imprimir 5
	x := sumHiguerOrder(1)
	y := x(3)
	fmt.Println(y) // Deve imprimir 4

	parcelas := descontoClosure(100)
	valorDaParcelaEm2x := parcelas(2)
	valorDaParcelaEm3x := parcelas(3)

	fmt.Println(valorDaParcelaEm2x) // Deve imprimir 45
	fmt.Println(valorDaParcelaEm3x) // Deve imprimir 30

	fmt.Println(listaDeNumeros(1, 2, 3, 4, 5)) // Deve imprimir 15

}

func somar(a int, b int) int {
	return a + b
}

// Quando os parâmetros são do mesmo tipo, pode-se omitir os tipos, deixando apenas o tipo no final
func subtrair(a, b int) int {
	return a - b
}

// A traducao de swap é "trocar"
// A função swap troca os valores de a e b
func swap(a, b int) (int, int) {
	return b, a // retornando dois valores
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	// return res, rem // res -> result | rem -> reminder -> resto
	// podemos usar o naked return, pois já declaramos as variáveis no retorno da função (res, rem)
	return // aqui ele vai retornar res e rem
	// o uso do naked return é desencorajado, pois pode tornar o código menos legível
}

// Funcoes Higher Order
func sumHiguerOrder(a int) func(int) int {
	// isso tambem é uma função closure, pois ela acessa a variável `a` que está no escopo da função pai `sumHiguerOrder`
	return func(b int) int {
		return a + b
	}
}

func descontoClosure(preco int) func(int) int {
	desconto := 10
	return func(parcelas int) int {
		return (preco - desconto) / parcelas
	}
}

func listaDeNumeros(nums ...int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	return total
}
