package main

import (
	"github.com/isaiaszmendes/gopportunities/config"
	"github.com/isaiaszmendes/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize the configuration
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing the configuration: %v", err)

		panic(err) // panic é uma função que interrompe a execução do programa
		// Nesse caso faz sentido pois se o banco de dados não estiver disponível, não faz sentido executar o programa
	}

	// Initialize the router
	router.Initialize()
}
