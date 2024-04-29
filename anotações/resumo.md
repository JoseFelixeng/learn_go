# Introcução a Linguagem golang

A extensão usada pela linguagem é o .go. 

O primeiro passo para que o arquivo seja executavél é a utilização do:

```go

package main  // usado para dizer onde esta o pacote a ser usado 

import (
     "fmt"
)
   

//Para o Go toda função precisa ter uma função man para executar os códigos 


func main(){

    fmt.Println("Olá, mundo!")

}


```

No go todos os programas são organizados em pacotes, um pacote é um grupo de arquivos que estão em um mesmo local e eles são compilados juntos. Dessa forma todos os arquivos serão compartilhados entre todos os pacotes. No programa acima podemos vê como se criar um código simples que mostre a mensagem olá mundo.

Um outro exemplo simples: 
```go
package main // usado para dizer onde esta o pacote a ser usado

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("my favorite Number is", rand.Intn(10))
}

```

### Variavéis no GO
Para declaramos a uma variavel podemos usar a palavra reserva da linguagem **var** e logo após o time de dado ao qual a variavel irá receber.

Exemplo: ``` nome var string ```

A instrução ```var``` pode ser usada no escopo de pacote ou de método/função, a instrução var pode incluir inicializadores, um por variável. Neste caso, não é necessário informar o tipo pois o mesmo será inferido.

```go 
package main

import(
	"fmt"
	"math"
)


var  c, python, java
func main (){
	var i int 
	fmt.Println(i, c, python, java)
}


```

* Dentro de uma função, o operador de declaração curta **:=** , pode ser usado ao invés do operador var, ao usar o operador de declaração curta o mesmo iniciará as variáveis

Digitando o comando ```go env``` é possível vê as variaveis de ambiente do go
O código  ```GOPATH: endereço go``` serve para guardar o binario implementados.

todos os ´´´packages´´´ usados no GO teram que possuir o mesmo nome da pasta onde o arquivo esta armazenado, tirando o ´´´main´´´ pois é onde estará as principais funções 

#### Declarando uma variavel 

```Go

package main 

const a = "Hello Wolrd"

var b bool

b = true 

func main(){
	println(b)
}

```

```Go 
//Declaração de Escobo Global 

var (
	c int
	b bool
)
```

Variaveis e Importes que não são usados o Go reclama e gera um erro.

É possível ainda declarar variaveis na forma de sort slice é usado para simplificar as declarações de variaveis.

```Go

package main 


func main(){
	a := "Felix" //Tipo String
	b := 23 // Tipo int

	println(a)
}
```

# Tipagem de dados em GO

```GO
package main

const a = "Hello Wolrd"
type ID 

var (
	c int     = 9
	b bool    = true
	d float64 = 2.3
	e string  = "Felix"
	f ID = 1 
)

func main() {
	b = true
	println(b)
	println(c)
}

```

O "fmt" é um pacote usado para formatação dos dados exibidos

```GO
package main

import "fmt"

const a = "Hello Wolrd"
type ID int 

var (
	c int     = 9
	b bool    = true
	d float64 = 2.3
	e string  = "Felix"
	f ID = 1 
)

func main() {

	fmt.print("Esse é o tipo de %t", f)
}

```