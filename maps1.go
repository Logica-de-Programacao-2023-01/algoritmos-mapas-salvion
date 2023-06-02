package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "a a a b c c"
	ocorrencias := caracteres(texto)
	fmt.Println(ocorrencias)
}

func caracteres(texto string) map[string]int {
	palavras := strings.Fields(texto)
	ocorrencias := make(map[string]int)
	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}
	return ocorrencias
}
