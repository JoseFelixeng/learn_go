package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Aqui é um texto\n"

	// Chamando as variaveis
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Println("O tipo da variavel é %t e o valor é %v\n", t, t)
}
