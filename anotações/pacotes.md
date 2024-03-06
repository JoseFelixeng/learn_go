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
    2. 
```
        package auxiliar

            import "fmt"

            func Escrever() {
                fmt.Println("escrevendo com outra função!")
            }
```


