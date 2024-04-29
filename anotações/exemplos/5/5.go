package main

import "fmt"

func main() {
	var meuArray [3]int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	//Percorrendo o Array

	for i, v := range meuArray {
		fmt.Printf("Percorrendo o array/vetor %d no indice com valor %d \n", i, v)
	}
}
