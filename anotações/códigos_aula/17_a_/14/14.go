package main

func soma(a, b *int) int {
	*a = 40
	*b = 80
	// Se os valores de a e b forem alterados aqui, os dados não são alterados nos valores das entradas
	// logo o ideal é passar a referencia da memoria na chamada da função e a função soma deve reveber ponteiros
	return *a + *b // Retornando o valor da memoria
}

func main() {
	m1 := 20
	m2 := 30
	// A função soma faz uma copia das variaveis
	// computa a ação e retorna o resultado
	println(soma(&m1, &m2))
	println(m1)
	println(m2)

}
