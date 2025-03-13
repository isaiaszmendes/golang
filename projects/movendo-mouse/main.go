package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	x := 3100
	y := 200

	for {
		robotgo.MoveSmooth(x, y, 5.0, 10.0)
		x += 50
		// gerando um valor aleatório para y entre dois valores
		y = rand.Intn((600-200)+1) + 200
		// printando o valor de x e y
		fmt.Println(x, y)
		if x > 3800 {
			x = 3100
		}

		time.Sleep(60 * time.Second)
	}
}
