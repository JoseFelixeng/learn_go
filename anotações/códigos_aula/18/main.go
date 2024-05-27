package main

//Constract
// Usado para agregar varios tipos
type Number interface {
	int | float64
}

type myNumber int

// Usando Generics
func SomaG[T Number](m map[string]T) T {
	var somaf T
	for _, v := range m {
		somaf += v
	}
	return somaf
}

func main() {
	m := map[string]int{"jose": 700, "Mateus": 1200, "Natalia": 10000}
	m1 := map[string]float64{"jose": 70.50, "Mateus": 12.99, "Natalia": 10.59}
	m2 := map[string]myNumber{"jose": 700, "Mateus": 1200, "Natalia": 10000}
	println(SomaG(m))
	println(SomaG(m1))
	println(m2)

}
