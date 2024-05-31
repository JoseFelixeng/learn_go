package main

import "fmt"

const a = "Hello Wolrd"

type ID int

var (
	c int     = 9
	b bool    = true
	d float64 = 2.3
	e string  = "Felix"
	f ID      = 1
)

func main() {

	fmt.Printf("Esse é o tipo de %t", f)
}
