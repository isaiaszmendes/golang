package main

// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// Demonstre estes valores.

func main() {
	const (
		_ = 2025 + iota
		year1
		year2
		year3
		year4
	)

	println(year1, year2, year3, year4)
}
