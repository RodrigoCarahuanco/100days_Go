package main

import "fmt"

func main() {
	array := []int{12, 43, 12, 32}
	resultado := numeroGrande(array)
	fmt.Println(resultado)
}

func numeroGrande(array []int) int {
	maximo := array[0]
	for _, num := range array {
		if num > maximo {
			maximo = num
		}
	}
	return maximo
}