package main

import "fmt"

func main() {
	var edad int

	fmt.Print("Ingrese su edad por favor: ")
	fmt.Scanln(&edad)

	if edad > 18 {
		fmt.Print("Ustede es mayor de 18")
	} else {
		fmt.Print("Usted no es mayor a 18")
	}

}