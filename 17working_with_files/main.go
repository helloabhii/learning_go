package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files ")
	content := "this is sample text and we will put this into a file"

	file, err := os.Create("./helloabhii.txt")

	if err != nil {
		panic(err) //panic shut down the execution of the program and shows what error you are facing

	}
	//to write into the file
	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("length is : ", length)
	defer file.Close() //defer is recommended here
	readFile("./helloabhii.txt")
}

///reading the file

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is  \n", string(databyte)) // here we turn datatype into string

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
