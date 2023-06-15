package main

import "fmt"

// Escreva uma função que receba uma lista de mapas, onde cada mapa contém a
// contagem de palavras de um texto, e retorne um único mapa contendo a soma de todas as contagens.

func juntar(listas []map[string]int) map[string]int {
	novo := make(map[string]int)
	for i := range listas {
		for chave, valor := range listas[i] {
			novo[chave] += valor
		}
	}
	return novo
}

func main() {
	lista1 := map[string]int{"maca": 10, "laranja": 20, "e": 40}
	lista2 := map[string]int{"boi": 10, "laranja": 10, "de": 3}
	lista3 := map[string]int{"meu": 10, "filho": 10, "de": 50}
	listas := []map[string]int{lista1, lista2, lista3}

	nova_lista := juntar(listas)
	fmt.Print(nova_lista)
}
