package main

import (
	"fmt"
)

type No struct {
	Chave    int
	Esquerda *No
	Direita  *No
}

type ArvoreBinaria struct {
	Raiz *No
}

func (arvore *ArvoreBinaria) Inserir(chave int) {
	novoNo := &No{Chave: chave}

	if arvore.Raiz == nil {
		arvore.Raiz = novoNo
	} else {
		inserirNo(arvore.Raiz, novoNo)
	}
}

func inserirNo(no, novoNo *No) {
	if novoNo.Chave < no.Chave {
		if no.Esquerda == nil {
			no.Esquerda = novoNo
		} else {
			inserirNo(no.Esquerda, novoNo)
		}
	} else if novoNo.Chave > no.Chave {
		if no.Direita == nil {
			no.Direita = novoNo
		} else {
			inserirNo(no.Direita, novoNo)
		}
	}
}

func (arvore *ArvoreBinaria) Buscar(chave int) bool {
	return buscarNo(arvore.Raiz, chave)
}

func buscarNo(no *No, chave int) bool {
	if no == nil {
		return false
	}

	if chave == no.Chave {
		return true
	} else if chave < no.Chave {
		return buscarNo(no.Esquerda, chave)
	} else {
		return buscarNo(no.Direita, chave)
	}
}

func main() {
	arvore := &ArvoreBinaria{}

	chaves := []int{10, 5, 15, 3, 7, 12, 17}

	for _, chave := range chaves {
		arvore.Inserir(chave)
	}

	fmt.Println("Inserção concluída!")

	chaveBusca := 7
	if arvore.Buscar(chaveBusca) {
		fmt.Printf("A chave %d foi encontrada na árvore.\n", chaveBusca)
	} else {
		fmt.Printf("A chave %d não foi encontrada na árvore.\n", chaveBusca)
	}
}
