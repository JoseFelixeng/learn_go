package matematica

import "fmt"

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() {
	fmt.Println("O carro está andando!")
}

func (c Carro) Andarr() string {
	return "O carro está andando!"
}
