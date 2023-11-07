package main

import "fmt"

func main() {
	greet()
	fmt.Println("Functions in golang !")

	result := add(3, 6) //add is the function name
	fmt.Println("Sum of two Numbers : ", result)

	////
	proadd := advadd(2, 5, 7, 9, 11)
	fmt.Println("Advance add result :", proadd)
}
func add(val1 int, val2 int) int {
	return val1 + val2
}

// /////
func advadd(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
func greet() {
	fmt.Println("Namste !!!")
}
