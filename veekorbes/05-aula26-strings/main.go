package main

import "fmt"

// Em Go, strings são imutáveis
// Se eu quiser alterar uma string, tenho que criar uma nova
func main() {
	str := "Hello, Go!"
	fmt.Printf("%v \n%T\n", str, str)

	// Convertendo string para slice de bytes
	sliceByte := []byte(str)

	fmt.Printf("=> %v \n=> %T\n", sliceByte, sliceByte)

	// Convertendo slice de bytes para string
	for _, v := range sliceByte {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
		// %v valor
		// %T tipo
		// %#U unicode
		// %#x hexadecimal
	}

	// Se quero uma formatação específica, uso a crase ``
	strFormated := `
	Hello, Go   !
	tudo 
	
	
	bem?`

	fmt.Printf("%v \n%T\n", strFormated, strFormated)

	phrase := "Hello, Isaías!"

	for _, v := range phrase {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(phrase); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", phrase[i], phrase[i], phrase[i], phrase[i])
		// No UTF-8, alguns caracteres ocupam mais de um byte
		// Que é o caso do "í" que ocupa 2 bytes
		// por isso, o valor fica errado nesse loop
	}

}
