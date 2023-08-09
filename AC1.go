package main

import "fmt"

func main(){
	fmt.Println(e_primo(7))
	fmt.Println(e_primo(20))

	fmt.Println(calculaFibo(5))
	fmt.Println(calculaFibo(7))

	fmt.Println(diaSemana(2))
	fmt.Println(diaSemana(4))

	fmt.Println(e_bissexto(2014))
	fmt.Println(e_bissexto(2016))
	
}


func e_primo(n int) bool{
	primo := true

	for i := 2; i <= n-1; i++{
		if n%i == 0{
			primo = false
			fmt.Println(i)
		}
	}
	return primo
}



func calculaFibo(n int) int{
	ult := 1
	penult := 1
	var fibo int

	for i := 3; i <= n; i++{
		fibo = ult + penult
		penult = ult
		ult = fibo
	}
	return fibo
}



func diaSemana(num int) string{
	switch num{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Valor inválido!"
	}
}


func e_bissexto(n int) bool{
	    if (n%4 == 0 && n%100 != 0) || n%400 == 0 {
        return true
    } else {
        return false
    }
}

