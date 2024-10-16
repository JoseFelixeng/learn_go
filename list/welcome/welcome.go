package main

import "fmt"

func Welcome(nome string) {
	fmt.Println("Welcome to party, %s \n ", nome)
}

func main() {
	var nome = "felix"
	Welcome(nome)
}
