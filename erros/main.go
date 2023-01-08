package main

import (
	"net/http"
	"log"
	"fmt"
	"errors"
)

func main() {
	res, err := http.Get("https://go.dev/")
	if err != nil {
		log.Fatal(err.Error())
	} 
	fmt.Println(res.Header)

	resp, erro := soma(10,2)
	if erro != nil {
		log.Fatal(erro.Error())
	} 
	fmt.Println(resp)

	// Quando não quero tratar erros, posso colocar o "_"
	// resposta, _ := soma(10,2)
	// fmt.Println(resposta)
}

func soma(x int, y int)(int, error) {
	res := x + y
	if res > 10  {
		return 0, errors.New("Total maior que 10");
	}

	return res, nil
}