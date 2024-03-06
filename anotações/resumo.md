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