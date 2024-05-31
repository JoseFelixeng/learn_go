// Aula sobre Struct
package main

import "fmt"

type Pessoa interface { // Exboço de uma interface em Go
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	// Estrutura de dados usada para um cliente qualquer
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Instanciando a estrutura endereço
}

func (c Cliente) Desativar() { // Dessa forma esta sendo criado um método que pode mudar o valor de uma struct
	// Criar um metodo para mudar um valor já ativo
	c.Ativo = false // Levar o cliente para False
	fmt.Printf("O usuario  %s foi Desativado %t.\n", c.Nome, c.Ativo)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	// Intanciando um Cliente
	felix := Cliente{
		Nome:  "Felix",
		Idade: 28,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}
	//felix.Cidade = "Santa Cruz" // usando struct como composição
	felix.Numero = 45 // usando struct como tipo

	Desativacao(minhaEmpresa)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Numero: %d", felix.Nome, felix.Idade, felix.Ativo, felix.Numero)

}
