package meupacote

// var baz is unused (U1000)go-staticcheck
// Isso acontece pois a variável baz não é exportada por não ser Publica, dai por enquanto ela não está sendo usada
var baz string = "Hello Baz"
