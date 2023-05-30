package main

import "fmt"

func main() {
	var confirmacion string

	fmt.Print("Confirme la compra Y/N: ")
	fmt.Scanln(&confirmacion)

	switch {
		case confirmacion == "y" || confirmacion == "Y":
			fmt.Print("Gracias por su compra")
		case confirmacion == "n" || confirmacion == "N":
			fmt.Print("Entendemos, vuelva pronto") 
	}

	switch confirmacion{
		case "y","Y":
			fmt.Print("Gracias por su compra")
		case "m","N":
			fmt.Print("Vuelva Pronto")
	}
}