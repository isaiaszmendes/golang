package main

import "fmt"

// variavel package-level scope
// possui scopo que abrange o package inteiro
var numeroComPackageLevelScope = 200

// se não atribuirmos um valor na declaração com o var
var m int

//! Isso da erro, só podemos atribuir dentro de um code block
//* m = 10

func main() {
	countBytes, erros := fmt.Println("Hello, world!", 100)

	// pacote.Indentificador
	fmt.Println(countBytes, erros)

	// underscore character
	_, err := fmt.Println("Hello, world!", 100)
	fmt.Println(err)

	// := serve para declarar variaveis
	// tipagem automatica
	// Só é acessivel dentro de escopos, code block {}, se quero declarar uma variavel fora do code block de {}
	// Tenho que usar o "var"

	x := 10
	y := "Isaias"
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v,  %T\n", y, y)

	z := 50
	qualquerNumero(z)

	printsFMT()
}

func qualquerNumero(x int) {
	fmt.Println(numeroComPackageLevelScope)
	fmt.Println(x)
}
func printsFMT() {
	x := 10
	fmt.Println(x)

	s := "bom"
	r := "dia"
	z := fmt.Sprint(s, r, 10, " ", "!") // Não imprime, apenas formata os valores e retorna uma string sem espaços
	fmt.Println(z)
	//* retorno: bomdia10 !

}
