package main

import "fmt"

type Contato struct {
	Nome string
	Email string
}

func (c *Contato) alteraEmail (email string){
	c.Email = email
	
}

func adicionaContato(c Contato, a [5]Contato) [5]Contato{
	for i, elemento := range a {
		if (elemento == Contato{}){
			a[i] = c
			break
		}
	} 
	return a
} 

func excluiContato(a [5]Contato)[5]Contato{
	for i, elemento := range a {
		if (elemento == Contato{}){
			a[i-1] = Contato{}
			break
		}
	}
	return a
}

func main(){
	for {
		switch escolha{
		case nome, email:
			var nome, email string
			fmt.Println("Nome: ")
			fmt.Scanln(&nome)
			fmt.Println("Email: ")
			fmt.Scanln(&email)
			
			c := Contato{
			Nome: nome,
			Email: email,
			}
		} 
	}

	


	c.alteraEmail("luispasturamacedo@gmail.com")
	fmt.Println(c)


}