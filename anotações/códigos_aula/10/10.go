// Aula Sobre Funções no GO

// Função main
package main

import (
	"fmt"
)

func main() {

	// podemos fazer também uma função que chame uma outra função dentro da propria main
	total := func() int { // a variavel total vai executar uma função e receber o valor de retorno
		return sum(10, 20, 30, 40, 50) * 2
	}() // forma correta de instanciar

	fmt.Println(total) // Chamando a  função soma.
}

func sum(numeros ...int) int { // usado para uma entrada de n itens
	total := 0                       // instanciando uma variavel pra ser usada como controle
	for _, numero := range numeros { // for entrada_não_instanciar, numero(nova_variavel) := range(usada percorrer todos os numeros)
		total += numero // total vai receber a atualização de numero
		fmt.Printf("valor atual %d \n", total)
	}
	return total // retorna a soma de todos os numeros
}
