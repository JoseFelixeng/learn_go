package main

//Biblioteca usada para trabalhar com arquivos no Go
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Inicialização ao usar arquivos no Go
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// Escrevendo em um arquivo
	//tamanho, err := f.WriteString("Hello,World")
	// escrevendo no arquivo outros dados que não são strings
	tamanho, err := f.Write([]byte("Escrevendo uma frase grande em um arquivo!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso!, tamanho: %d bytes", tamanho)
	f.Close()

	// Abrindo um arquivo para leitura
	// os.Open() => abri o arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// Usando o bufio para criar buffer para leitura do arquivo
	// transformar o codigo em pedaços e reconstroi usando um slice
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
