package main

import (
	"fmt"
	"strconv"
)

//Escreva uma função que receba um slice de inteiros e retorne um mapa onde as chaves
//são os pares de números encontrados no slice e os valores são as frequências de cada par.

func FreqPar(num []int) map[int]int {
	mapaFreq := make(map[int]int)

	for i := 0; i < len(num)-1; i++ {
		pares := strconv.Itoa(num[i]) + strconv.Itoa(num[i+1])
		par, err := strconv.Atoi(pares)
		if err != nil {
			fmt.Println("Erro ao converter a string para int:", err)
			return nil
		}
		mapaFreq[int(par)]++
	}
	return mapaFreq
}

func main() {
	numeros := []int{1, 2, 6, 4, 3, 3, 2, 1, 2, 6, 4, 4, 3}
	frequencia := FreqPar(numeros)
	for par, freq := range frequencia {
		fmt.Printf("Par: %d, frequencia: %d\n", par, freq)
	}
}
