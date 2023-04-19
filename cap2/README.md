Tipos em go são estáticos

:= serve para declarar variaveis
tipagem automatica
Só é acessivel dentro de escopos, code block {}, se quero declarar uma variavel fora do code block de {}
Tenho que usar o "var"

Quando declaramos usando o var, se não atribuirmos um valor na mesma linha, só podemos atribuir dentro de um code block

var x declaração
y := 1 declaração, inicialização e atribuição
y = atribuição

x = 50 inicialização e atribuição

### Os Zeros

- int: 0
- floats: 0.0
- booleans: false
- strings: ""
- pointers, functions, interfaces, slices, channels, maps: nil

Use := sempre que possível
Use var para package-level scope