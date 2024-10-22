package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Podemos usar tgs nas struct

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta) // marshal usado par transformar o dado em byte.
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res)) // converto o binario json em string

	err = json.NewEncoder(os.Stdout).Encode(conta) // faz a serialização do dado

	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n": 2 , "s": 200}`)
	var contaX Conta
	// fazendo o teste contrario
	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)

}
