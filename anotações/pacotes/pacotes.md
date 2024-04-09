# Pacotes do Go 

1. Cada programa em Go é contido em um pacote
2. Programas iniciam a partir do pacote main
3. No go quando se esta utilizando mais de um pacotes é necessario criar um modulo
4. O moculo é um conjunto de pacotes que compõem o projeto.
5. Para criar o modulo é bem simples 
    1. É só abrir a pasta ao qual o arquivo será executado 
    2. E executar o comando ``` go mod init <nome_modulo> ``` 
    3. Após isso será criado um arquivo de nome **<nome_modulo>.mod**
    4. Dentro do arquivo encontraremos as seguintes informaçõs: 
        ```
        module modulo

        go 1.22.1

        ```
Dessa forma podemos vê que uma arquivo é criado com o nome do modulo desejado e também a versão do go que estamos usando. Para criar um modulo é interessante criar uma nova pasta e nessa pasta colocar o nome do modulo ao qual se deseja criar.

### Exemplo: 
1. ``` cd /documents/learn_go/```
2. ``` mkdir teste ```
3. logo após criamos  um arquivo dentro da pasta ``` teste ```
    1. ```teste.go```

```go
        package auxiliar

            import "fmt"

            func Escrever() {
                fmt.Println("escrevendo com outra função!")
            }
```

Usando o comando go biuld ele vai criar um arquivo ```go.mod```, é um arquivo binario executável no qual ele compila todo o projetos. Para rodar o arquivo compilado é necessario apenas usar o comando "./main" sendo **main** o arquvo atual que estamos usando.

    - Todo novo pacote deve esta contido em uma pasta separada.
    - Para importa um modulo é necessario main\auxiliar

Exemplo: 

```go
package main

import (
	"fmt"
	"main/auxiliar"
)

func main() {
    // Função usada para escrever
	fmt.Println("escrevendo do arquivo main")
	auxiliar.Escrever()
}
````
Como o Go não descreve se um arquivo é "public" ou "private", se a função estiver escrita com a primeira letra minuscula a mesma só será reconhecida dentro do pacote que a usa. já se ela começar com letra maiuscula o mesmo pode ser importado por outros pacotes, no Go se uma função é exportada a mesma deve conter um comentario.

O arquivo de modulo não se atualiza sozinho logo o mesmo sempre deve esta sendo atualizado para evitar erros ou maiores complicações.
O comando go install vai usar  o pacote criado e o levará para a pasta de instalação do Go dessa forma o mesmo poderá ser usado como um pacote padrão da linguagem.