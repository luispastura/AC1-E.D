package main

import (
	"ac/contatos"
	"ac/oper"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var ctts [5]contatos.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			ctts = oper.AdicionaContato(contatos.Contato{Nome: nome, Email: email}, ctts)
		case "2":
			ctts = oper.ExcluiContato(ctts)
		default:
			return
		}

		fmt.Println(ctts)
	}

}



	
	
