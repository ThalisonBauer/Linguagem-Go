package main
import  "fmt"
func FUNC1(){
	 var B int
	B=-100
	fmt.Println("Valor de B dentro da função FUNC1:",B)
}
func FUNC2()  {
	var B int
	B=-200
	fmt.Println("Valor de B dentro da função FUNC2:",B)
}
func main(){
	var B int
	B=10
	fmt.Println("Valor de B:",B)
	B=20
	FUNC1()
	fmt.Println("Valor de B:",B)
	B=30
	FUNC2()
	fmt.Println("Valor de B:",B)
}
