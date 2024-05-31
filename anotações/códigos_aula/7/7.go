package main

import "fmt"

func main() {

	salario := map[string]int{"Felix": 700, "Tadeu": 3200, "Natalia": 20000}

	fmt.Println("O salario do Funcionario é: ", salario["Felix"])

	delete(salario, "Felix")
	salario["felix"] = 1000

	fmt.Println("Apos ser contratado o salario do Funcionario é: ", salario["felix"])

	sal := make(map[string]int)

	sal["Felix"] = 1000

	for nome, salario := range salario {
		fmt.Printf("O salario é de %s e %d\n", nome, salario)
	}
}
