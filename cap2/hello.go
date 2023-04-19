package main

import "fmt"
// variavel Package Level Scope
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
}

func qualquerNumero(x int) {
	fmt.Println(numeroComPackageLevelScope)
	fmt.Println(x)
}
