package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	fmt.Printf("O cliente %v andou", c.nome)
}

func main() {
	wesley := Cliente{
		nome: "Wesley",
	}
	wesley.andou()

}
