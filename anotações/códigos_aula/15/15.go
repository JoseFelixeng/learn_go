package main

import "fmt"

type Conta struct {
	saldo float64
}

// Criando uma função que inicia com valor zerado
// Atuando como um construtor
func NewConta() *Conta {
	return &Conta{saldo: 0} // Criando uma conta com saldo zerado.
}

// Função para simular o saldo
func (c *Conta) simular(valor float64) float64 {
	// Ao usar ponteiro é armazenado direto no lugar da memoria o valor de conta
	// Nesse caso o valor será alterado diretamente na memoria
	c.saldo += valor
	fmt.Println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	fmt.Println(conta.saldo)

}
