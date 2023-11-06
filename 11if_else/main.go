package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 21
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "exactly equal to 10 login count"
	}

	fmt.Println(result)

	/// practice

	var2 := 71
	var ans string
	if var2 < 21 {
		ans = "variable is smaller"
	} else if var2 > 71 {
		ans = "variable is larger"
	} else {
		ans = "variable is equal to 71"
	}
	fmt.Println("Result : ", ans)

	//ques guess colour
	color := "Pink"
	var clr string
	if color == "Pink" {
		clr = "Right Guess !!! Pink "
	} else if color == "Blue" {
		clr = "Right Guess !!! Blue "
	} else if color == "Yellow" {
		clr = "Right Guess !!! Yellow "
	} else {
		clr = "Color Not Found !"
	}
	fmt.Println(clr)

	//to get even or odd number
	if 7%2 == 0 {
		fmt.Println("Number is Even !")
	} else {
		fmt.Println("Number is Odd ! ")
	}
	////
	if num := 7; num < 7 {
		fmt.Println("Num is Less than 7")
	} else if num > 7 {
		fmt.Println("Num is Greater than 7")
	} else {
		fmt.Println("Num is equal to 7 ")
	}

	/// if err

	//	if err != nil {
	//		fmt.Println("Error Found")
	//	}
}
