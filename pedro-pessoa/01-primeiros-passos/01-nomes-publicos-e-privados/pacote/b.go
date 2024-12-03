package pacote

import (
	"fmt"
	"myFirstGoModule/pacote/internal/foo"
)

var Bar string = "Hello Bar"

func PrintMinha() {
	fmt.Println(foo.Minha)
}
