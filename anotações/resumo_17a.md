# Funções Variadicas

São funções que vem implementadas junto a propria linguagem GO e podem ser usadas pelos programadores em suas atividades.

```GO


func sum(numetos ...int) int { // Essa forma é usada quando não sabemos a qualidade de dados da entrada 

}
```

Dessa forma pode ser feito um Loop dentro dos parâmentros usados, de forma a usar uma infinidade de dados da entrada.

```Go

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

```
Exemplo usando alguma das informações das aulas anteriores. 


**range** pode ser usada para percorrer um dado elemento, ou parametro de entrada.

Podemos fazer uma função que possa ser chamada dentro da propria main 

```Go
func main() {

	// podemos fazer também uma função que chame uma outra função dentro da propria main
	total := func() int {
		return sum(10, 20, 30, 40, 50) * 2
	}() // forma correta de instanciar

	fmt.Println(sum(3, 4, 5, 6, 7, 10)) // Chamando a  função soma.
}

```

Exemplo: 

```Go 

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

```


# Structs 

Obs: O go não é uma linguagem orientada a objetos 

O struct no Go seria o mais proximo na linguagem de uma `class` em outra linguagem, mas lembrando uma struct não é uma classe.

```go 
// Aula sobre Struct
package main

import "fmt"

type Cliente struct {
	// Estrutura de dados usada para um cliente qualquer
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	// Intanciando um Cliente
	felix := Cliente{
		Nome:  "Felix",
		Idade: 28,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", felix.Nome, felix.Idade, felix.Ativo)

}

```

No Go não é possível usar herança, mas é possível usar composição para fazer o mesmo papel. Uma struct pode ser usada como um tipo dentro de outra, dessa forma além da composição poderá ser usada como um tipo dentro de outra struct.


**Usando como tipo**

```GO
type Cliente struct {
	// Estrutura de dados usada para um cliente qualquer
	Nome   string
	Idade  int
	Ativo  bool
	Adress Endereco // Instanciando a estrutura endereço
}

```

**Usando como Composição**

```GO
type Cliente struct {
	// Estrutura de dados usada para um cliente qualquer
	Nome   string
	Idade  int
	Ativo  bool
	Endereco // Instanciando a estrutura endereço
}

``` 


#### Exemplo

```GO

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

``` 

## Métodos de Struct 

Uma Struct também possui metodos.

```GO

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

func (c Cliente) Desativar() { // Dessa forma esta sendo criado um método que pode mudar o valor de uma struct
	// Criar um metodo para mudar um valor já ativo
	c.Ativo = false // Levar o cliente para False
	fmt.Printf("O usuario  %s foi Desativado %t.\n", c.Nome, c.Ativo)
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

	felix.Desativar() // chamando o metodo criado

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Numero: %d", felix.Nome, felix.Idade, felix.Ativo, felix.Adress.Numero)

}
```

A cima podemos vê um exemplo de utilização de uma  função para alterar uma struct de maneira semelhante ao que seria um metodo de uma class.


# Interfaces 

No caso do Go não precisa ser feito uma atachment para a implemetação de uma interface em uma estrutura, nos possibilitando utilizar diversos tipos de uma forma mais simples abstraindo muitas informações na hora implementação. De forma simples qualquer  metodo criado, caso o mesmo possua uma interface o mesmo poderá ser implementado, veja no exemplo a seguir.

```GO

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
```

Uma observação importante é que a interface no GO permite apenas a utilização de metodos feitos, logo declarações de variaveis e outros exemplos, não podem ser declarados em uma interface no GO.

# Ponteiros