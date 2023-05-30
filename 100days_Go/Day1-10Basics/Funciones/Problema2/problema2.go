package main

import "fmt"

func main() {
	problema2()
}

func problema2() {
	var dividendo int = 12
	var divisor int = 6

	var cociente = dividendo / divisor
	var residuo = dividendo % divisor

	fmt.Println("Cociente: ", cociente)
	fmt.Println("Residuo: ", residuo)
}