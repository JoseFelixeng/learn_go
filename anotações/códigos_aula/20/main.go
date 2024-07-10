package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}
	// um For menos convencional

	for k, v := range numeros {
		// nesse caso
		// k indica a posição do array
		// v indica o valor atribuido a posição que nesse caso é uma strings
		println(k, v)
	}

	for {
		println("Hello,World")
	}
}
