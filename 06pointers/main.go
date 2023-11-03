package main

import (
	"fmt"
)

func main() {
	fmt.Println("Introduction to Pointers !!!") // passing the memory address
	//pointer is the reference to the memory addres //means actual is passsed on
	var one *string //means that this data type var is pointer //one is the pointer here
	fmt.Println("Value of Pointer is : ", one)

	anyInt := 71
	var ref = &anyInt                            // reference to the address
	fmt.Println("Value of the pointer : ", ref)  // memory location
	fmt.Println("Value of the pointer : ", *ref) //value of actual pointer

	*ref = *ref - 2 // pointer value change means actual value also change
	fmt.Println("New value :", anyInt)
}
