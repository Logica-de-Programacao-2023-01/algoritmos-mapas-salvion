package main

import "fmt"

//Escreva uma função que gere a sequência de Fibonacci até um determinado número n e
//retorne um mapa onde as chaves são os índices da sequência e os valores são os números correspondentes.

func fibonacci(n int) map[int]int {
	mapa := make(map[int]int)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			mapa[i] = 1
		} else {
			mapa[i] = mapa[i-1] + mapa[i-2]
		}
	}
	return mapa
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Print(fibonacci(n))
}
