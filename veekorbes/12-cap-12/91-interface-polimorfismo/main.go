package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesArrancados string
	salario          float64
}

type arquiteto struct {
	pessoa
	tipoConstrucao string
	tamanhoProjeto float64
}

func (x dentista) oiBomDia() {
	fmt.Printf("Oi, bom dia! Meu nome é %s e arranquei %s dentes\n", x.nome, x.dentesArrancados)
}

func (x arquiteto) oiBomDia() {
	fmt.Printf("Oi, bom dia! Meu nome é %s %s, sou arquiteto\n", x.nome, x.sobrenome)
}

type gente interface {
	oiBomDia()
}

func serhumano(g gente) {
	// g.oiBomDia()
	switch g.(type) {
	case dentista:
		fmt.Printf("Sou um dentista e ganho %.2f\n", g.(dentista).salario)
	case arquiteto:
		fmt.Printf("Sou um arquiteto e meu projeto tem %.2f m²\n", g.(arquiteto).tamanhoProjeto)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Dr. Dente",
			sobrenome: "Silva",
			idade:     45,
		},
		dentesArrancados: "10",
		salario:          10000.00,
	}

	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Eng.R",
			sobrenome: "Predio",
			idade:     35,
		},
		tipoConstrucao: "Prédio",
		tamanhoProjeto: 1000.00,
	}

	mrdente.oiBomDia()

	mrpredio.oiBomDia()

	serhumano(mrdente)
	serhumano(mrpredio)

}
