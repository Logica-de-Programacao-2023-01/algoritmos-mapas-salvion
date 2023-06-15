package main

import (
	"fmt"
)

// Escreva uma função que receba uma string e retorne um mapa onde as
// chaves são os caracteres presentes na string e os valores são a frequência de cada caractere.

func mapear(s string) map[string]int {
	mapa := make(map[string]int)
	for _, char := range s {
		mapa[string(char)]++
	}
	return mapa
}

func main() {
	s := "palavrona"
	mapa := mapear(s)
	fmt.Print(mapa)
}
