// Aula sobre Struct
package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	// Estrutura de dados usada para um cliente qualquer
	Nome   string
	Idade  int
	Ativo  bool
	Adress Endereco // Instanciando a estrutura endereço
}

func main() {
	// Intanciando um Cliente
	felix := Cliente{
		Nome:  "Felix",
		Idade: 28,
		Ativo: true,
	}

	//felix.Cidade = "Santa Cruz" // usando struct como composição
	felix.Adress.Numero = 45 // usando struct como tipo

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Numero: %d", felix.Nome, felix.Idade, felix.Ativo, felix.Adress.Numero)

}
