package main

import "fmt"

//Escreva uma função que receba um mapa onde as chaves são nomes de pessoas e os valores são as
//quantias de dinheiro que cada pessoa gastou em uma conta compartilhada.
//A função deve calcular o valor que cada pessoa deve receber ou pagar para igualar as despesas.

func igualar(gastos map[string]float64) {

	media := 0.0
	for _, valor := range gastos {
		media += valor
	}

	media = media / float64(len(gastos))
	for chave, valor := range gastos {
		resultante := valor - media
		if resultante > 0 {
			fmt.Printf("%s deve pagar %.2f.\n", chave, resultante)
		} else {
			resultante *= -1
			fmt.Printf("%s deve receber %.2f.\n", chave, resultante)
		}
	}
}

func main() {
	gastos_casa := map[string]float64{
		"Pedro": 23.50,
		"Clara": 100.50,
		"Ana":   87.99,
		"Lucas": 280.00,
	}
	igualar(gastos_casa)
}
