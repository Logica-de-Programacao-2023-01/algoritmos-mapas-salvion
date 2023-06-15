package main

import (
	"fmt"
	"strings"
)

//Escreva uma função que receba uma string contendo uma frase e retorne um mapa onde
//as chaves são as palavras encontradas na frase e os valores são mapas contendo a contagem
//de cada letra em cada palavra.

func freq_letras(s string) map[string]int {
	fletra := make(map[string]int)
	for _, letra := range s {
		fletra[string(letra)]++
	}
	return fletra
}

func freq_palavra(s string) map[string]map[string]int {
	//ex: hoje[map[letra]frequencia]

	mapa := make(map[string]map[string]int)
	s = strings.ToLower(s)
	S := strings.Fields(s)

	for _, palavra := range S {
		mapa[palavra] = freq_letras(palavra)
	}
	return mapa
}

func main() {
	frase := "Oi hoje é um belissimo dia, hoje é tudo que sonhei."
	mapa := freq_palavra(frase)
	fmt.Print(mapa)
}
