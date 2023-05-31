package main

import "fmt"

func main() {
	array := []int{1, 2, 3}
	resultado := sumaArray(array)
	fmt.Println(resultado)
}

func sumaArray(arr []int) int {
	suma := 0

	for _, num := range arr {
		suma += num
	}
	return suma
}