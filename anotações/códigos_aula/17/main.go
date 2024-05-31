package main

import "fmt"

func main() {
	var minhaVar interface{} = "Jose Felix"
	fmt.Println(minhaVar)
	println(minhaVar.(string))
	// Usado para mudar o tipo, se ok==true então a conversão pode ser feita
	res, ok := minhaVar.(int)

	fmt.Printf("O valor de res %v e ok é %v", res, ok)
}
