package main

import "fmt"

func encontraParComSoma(array []int, alvo int) (int, int) {
	numerosVistos := make(map[int]int)

	for indice, numero := range array {
		diferenca := alvo - numero

		if outroIndice, ok := numerosVistos[diferenca]; ok {
			return outroIndice, indice
		}

		numerosVistos[numero] = indice
	}

	return -1, -1
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	alvo := 7
	indice1, indice2 := encontraParComSoma(array, alvo)
	if indice1 != -1 && indice2 != -1 {
		fmt.Printf("Par encontrado: [%d, %d]\n", array[indice1], array[indice2])
	} else {
		fmt.Println("Nenhum par encontrado.")
	}
}
