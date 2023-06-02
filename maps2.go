package main

import "fmt"

//Escreva uma função que receba dois mapas e retorne um novo mapa contendo todos os elementos dos mapas de entrada.
//Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.

func main() {
	mapa1 := map[string]int{"mel": 1, "iogurte": 3, "granola": 30}
	mapa2 := map[string]int{"mel": 4, "banana": 3, "maça": 1}
	novoMapa := UnirMapa(mapa1, mapa2)
	fmt.Println(novoMapa)
}

func UnirMapa(m1, m2 map[string]int) map[string]int {
	mapa := make(map[string]int)
	for chave, valor := range m1 {
		mapa[chave] = valor
	}
	for chave, valor := range m2 {
		mapa[chave] = valor
	}
	return mapa
}
