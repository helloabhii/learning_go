package main

import "fmt"

const OutsideVar string = "declared outside any function" // you can use it anywhere in the file //public

func main() {
	fmt.Println("Variables")
	var name string = "helloabhii"
	fmt.Printf("Variable is of type : %T \n", name) // %T for type checking

	var rightorwrong bool = false                           //returns true or false only
	fmt.Printf("Variable is of type : %T \n", rightorwrong) // \n - used ofr new line or break the line

	var num int = 71
	fmt.Printf("Variable is of type : %T \n", num) // integer

	var flt float32 = 71.99 //float32 and float64 are present in the golang
	fmt.Printf("Variable is of type : %T \n", flt)

	// default valus and some aliases
	var anotherVar int //only decalring it // when we just declared the variable its value is 0 initially
	fmt.Println("decalared var : ", anotherVar)
	fmt.Printf("Var is type of : %T\n", anotherVar) // check its type // \n gives the extra space

	/////////////////////////////////////////////////////////////////////////////////////////////
	//implicit type
	/////////////////////////////////////////////////////////////////////////////////////////////

	var abhi = "any_string" // var if you again change it to int it would not be allowed it cause an error
	abhi = "myName"         //cannot cause error because you manipulate it to string
	// abhi = 71 // cause error because you change string to int
	fmt.Println(abhi)

	// ///////////////////////////////////////////////////////////////////////////////////////////
	// no var type // also called as short hand notation maybe
	// ///////////////////////////////////////////////////////////////////////////////////////////
	vr := 71 // once you create a var you have to use it
	fmt.Println("Var int :", vr)
	//outside method / you are not allowed to declare the va := method/ means short hand not allowed outside the function/method

	fmt.Println("Var is declare outside :", OutsideVar)
	fmt.Printf("This outside var is type of : %T", OutsideVar) // printf for format get out
}
