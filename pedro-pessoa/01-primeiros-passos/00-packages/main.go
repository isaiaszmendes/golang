package main

import (
	myPackageRenamed "encoding/json" // o ultimo nome da importacao é o que será usado no código, a menos que vc renomeie
	"fmt"
	// . "fmt" // o ponto permite que vc use os métodos do pacote sem precisar chamar o pacote
	// ex: Println("Hello World") ao invés de fmt.Println("Hello World")
	// mas não é recomendado, pois pode causar confusão
	// temos o erro Should not use dot imports (ST1001) no linter
	// _ "fmt" // professor disse que vai explicar isso depois
)

func main() {
	fmt.Println("Hello World")
	myPackageRenamed.Marshal()
}
