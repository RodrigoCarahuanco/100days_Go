package main

import "fmt"

func main() {
	problema3()
}

func problema3() {
	var monto float32

	fmt.Println("Ingresa el monto: ")
	fmt.Scan(&monto)

	var igv float32 = monto * 0.12
	var total float32 = monto - igv
	fmt.Println("IGV: ", igv)
	fmt.Println("Total: ", total)
}	