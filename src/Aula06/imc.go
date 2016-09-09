package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var soma float64
	var sexo int

	fmt.Println("Digite seu Peso:")
	fmt.Scanln(&peso)

	fmt.Println("Digite sua Altura:")
	fmt.Scanln(&altura)

	fmt.Println("1-MASCULINO 2-FEMININO")
	fmt.Scanln(&sexo)
	soma = peso/(altura*altura)
	switch sexo {
	case 1:
		if soma < 20.7{
			fmt.Println("Abaixo do peso Ideal")
		}else if soma > 20.7 && soma < 26.4{
			fmt.Println("No peso Ideal")

		}else if soma > 26.4 && soma < 27.8{
			fmt.Println("Marginalmento acima do peso")

		}else if soma > 27.8 && soma < 31.1{
			fmt.Println("Acima do peso ideal")
		}else{
			fmt.Println("OBESO")
		}
		break
	case 2:
		if soma < 19.1{
			fmt.Println("Abaixo do peso Ideal")
		}else if soma > 19.1 && soma < 25.8{
			fmt.Println("No peso Ideal")

		}else if soma >25.8 && soma < 27.3{
			fmt.Println("Marginalmento acima do peso")

		}else if soma > 27.3 && soma < 32.3{
			fmt.Println("Acima do peso ideal")
		}else{
			fmt.Println("OBESO")
		}
		break

	}

}
