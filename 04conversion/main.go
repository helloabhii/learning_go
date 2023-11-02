package main

import (
	"bufio" // for input purposes or other env var etc
	"fmt"
	"os"
	"strconv" //converting string into the numbers
	"strings"
)

func main() {
	/////////////////////////////////////////////////////////////////////////////////////////
	//Converstion
	/////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Choose No. between 1 -  5 ")
	reader := bufio.NewReader(os.Stdin)
	// to create a new reader that reads input from the standard input (the keyboard).
	input, _ := reader.ReadString('\n') //similar as try and except block in py
	// It reads the user input until the Enter key is pressed and stores it in the input variable as a string.
	fmt.Println("Thanks for writing : ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // function is used to convert the user's input (which is a string) to a floating-point number (with a precision of 64 bits).
	// is used to remove leading and trailing whitespace from the user's input to ensure that the conversion works correctly.
	// The result of the conversion is stored in the numRating variable, and any error that occurs during the conversion is stored in the err variable.
	if err != nil {
		fmt.Println(err) // if error occured stored in this variable
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
