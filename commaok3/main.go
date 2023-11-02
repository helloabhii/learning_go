package main

import (
	"bufio" // used for buffered I/O operations, such as reading from or writing to files and other data sources.
	"fmt"
	"os" // to get working with files, accessing data manipulation etc
)

func main() {
	in := "Taking input from user"
	fmt.Println(in)

	reader := bufio.NewReader(os.Stdin) // to get input from standard input(console) // reading from input(os)
	fmt.Print("Write anything : ")

	//comma ok /error ok

	input, _ := reader.ReadString('\n')             //  to read a line of text input until the Enter (newline) key is  pressed
	fmt.Println("Thanks for writing  :", input)     // here whatever  you input printed out
	fmt.Printf("Type of this  thing : %T\n", input) // get the type of it
}
