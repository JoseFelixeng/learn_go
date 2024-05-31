package main

import (
	"fmt"
	"matematica/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "fiat"}
	carro.Andar()
	fmt.Println("chamando a função para ", carro.Andarr())
	fmt.Println("O Resultaod : $", s)
	fmt.Println("A marca do carro é ", carro.Marca)
	fmt.Println("O valor de a na função matematica", matematica.A)
}
