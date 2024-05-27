// Aula Sobre Funções no GO

// Função main
package main

import (
	"errors"
	"fmt"
)

func main() {

	//Chamando a função em GO
	//a := sum(1, 2)

	//fmt.Printf("Func Soma = %d \n", a)
	// Imprimindo e fazendo a chamada

	//fmt.Println(sum(105, 2))

	//Criando uma nova variavel para armazenar o valor e outra para o error
	valor, err := sum(51, 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)

}

// Primeira função em Go
//func sum(a, b int) (int, bool) {
//	if a+b >= 50 {
//		return a + b, true
//	}

//	return a + b, false
//}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, nil // O erro aqui é vazio logo não existe
	}

	return 0, errors.New("A soma é menor que 50")
}
