package main

import "fmt"

func main() {

	s := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0}

	fmt.Printf("Ignora tudo apos : len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("Ignora apos os 4 primeiros len=%d cap=%d %v \n", len(s[4:]), cap(s[:4]), s[:4])

	fmt.Printf("Ignora os 2 primeiros len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)

	fmt.Printf("ADD e Dobrou len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])
}
