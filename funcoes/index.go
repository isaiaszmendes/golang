package main
import "fmt"

func main() {
	result := Soma(1,1)

	fmt.Printf("%T \n", result)
}

func Soma(a int, b int) int{
	return a + b
}