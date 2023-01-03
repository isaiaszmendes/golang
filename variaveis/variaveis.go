package main
import "fmt"
var nome string

func main() {
	nome = "isaias"
	x := "mendes"
	x= "Isaias F Mendes"
	fmt.Println(nome,x)

	a:= 10
	b:= "World"
	c:= 3.14
	d:= false
	e:= `Hi,


	man!
	`
	// Mostra o valor
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	
	// Mostra o type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)


}