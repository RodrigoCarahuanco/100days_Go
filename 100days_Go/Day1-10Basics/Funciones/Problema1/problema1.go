package main

import "fmt"

func main() {
	suma()
}

func suma() {
	var a int
	var b int
	fmt.Println("Vamos a sumar dos numeros: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	var suma int = a + b 
	fmt.Print("La suma total de las variables es: ",suma)
}