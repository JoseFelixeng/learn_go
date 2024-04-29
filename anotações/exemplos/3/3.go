package main

const a = "Hello Wolrd"

type ID int // type nome tipo_de_variavel

var (
	c int     = 9
	b bool    = true
	d float64 = 2.3
	e string  = "Felix"
	f ID      = 1
)

func main() {
	b = true
	println(b)
	println(c)
	println(f)
}
