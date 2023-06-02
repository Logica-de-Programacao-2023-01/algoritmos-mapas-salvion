package main

import "fmt"

//Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.

func main() {
	numeros := map[int]int{1: 2, 2: 3, 3: 9}
	soma := soma(numeros)
	fmt.Println(numeros, soma)
}

func soma(mapa map[int]int) int {
	soma := 0
	for _, valor := range mapa {
		soma += valor
	}
	return soma
}
