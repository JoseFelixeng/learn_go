// Aula Sobre Funções no GO

// Função main
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(3, 4, 5, 6, 7, 10)) // Chamando a  função soma.
}

func sum(numeros ...int) int { // usado para uma entrada de n itens
	total := 0                       // instanciando uma variavel pra ser usada como controle
	for _, numero := range numeros { // for entrada_não_instanciar, numero(nova_variavel) := range(usada percorrer todos os numeros)
		total += numero // total vai receber a atualização de numero
	}
	return total // retorna a soma de todos os numeros
}
