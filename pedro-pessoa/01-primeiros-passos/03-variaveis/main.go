package main

import "fmt"

// aqui não vai dar erro de variável não utilizada
var idade int

//! não pode usar := fora de funções
// nome := "isaias"

func main() {
	// pode ter escopo de pacote ou de função
	var nome, sobrenome string
	var peso float64

	// Em Go, todas as variáveis começam com valor zerados
	// string = ""
	// int = 0
	// float = 0
	// bool = false

	fmt.Println(nome, sobrenome, idade, peso)
	fmt.Printf("= %q = %q = %d = %.2f\n", nome, sobrenome, idade, peso)

	// declaração e atribuição
	var name, lastName string = "Isaias", "Mendes"

	fmt.Println(name, lastName)

	// declaração agrupada
	var (
		altura float64 = 1.80
		casado bool    = true
	)

	fmt.Println(altura, casado)

	// := só pode ser usado dentro do escopo de funções
	//! Quando for usar o :=, não posso declarar o tipo da variável
	vaiGravarProCanalDoYoutube := true
	nomeDoCanal := "captalento"
	//! Tem que tomar cuidado quando vc precisa de outro tipo de variável
	// como por exemplo, um int ou float, int != int8 != int16 != int32 != int64

	fmt.Println(vaiGravarProCanalDoYoutube, nomeDoCanal)

	// Posso declarar variáveis de tipos diferentes na mesma linha com var ou :=

	var foo, bar = "foo", 34
	fizz, buzz := "fizz", 42

	fmt.Println(foo, bar, fizz, buzz)

	//! Não posso declarar variáveis de tipos diferentes na mesma linha passando apenas os tipos
	// var inteiro, str int, string
}
