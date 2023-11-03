package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in Golang !!!")

	var arr [7]string // here you have to define explicitly arr of 7 items
	arr[0] = "Abhi"
	arr[1] = "Linux"
	fmt.Println(arr[0])
	fmt.Println(arr[3])
	fmt.Println(arr)
	fmt.Println(len(arr)) // shows the value which  you declare while creating array

	var anotherArr = [5]string{"Abhi", "Cloud", "DevOps"}
	fmt.Println("Another array List : ", anotherArr)
}
