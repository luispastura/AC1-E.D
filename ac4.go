package main

import (
	"fmt"
	"strconv"
)

func hanoi(n int, origem, destino, meio string) int {
	if n == 1 {
		fmt.Println("Mexa o disco 1 de " + origem + " para " + destino)
		return 1
	}

	movimentos := hanoi(n-1, origem, meio, destino)
	fmt.Println("Mexa o disco ", n, "de", origem, " para ", destino)
	movimentos++
	movimentos += hanoi(n-1, meio, destino, origem)

	return movimentos
}

func main() {
	n := 3
	origem := "A"
	destino := "C"
	meio := "B"

	movimentos := hanoi(n, origem, destino, meio)
	fmt.Println("-----------------------------------")
	fmt.Println("Total de movimentos a serem feitos: " + strconv.Itoa(movimentos))
}
