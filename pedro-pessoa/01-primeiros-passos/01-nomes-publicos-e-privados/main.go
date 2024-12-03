package main

import (
	"fmt"
	"myFirstGoModule/pacote"
	"myFirstGoModule/pacote/internal/foo" // em Go, quando eu crio um diretorio chamado internal,
	// eu nao consigo acessar o pacote de fora do pacote em que internal esta
	// ex
	// pacoteA/internal/foo
	// pacoteA/print
	// pacoteB/hello
	// main
	// foo so pode ser acessado dentro de pacoteA, nao pode ser acessado por pacoteB, muito menos por main 

func main() {
	fmt.Println("Hello World")
	pacote.PrintMinha()
	foo.Minha
}
