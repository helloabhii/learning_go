package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to math in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4

	// fmt.Println("The sum is : ", mynumberOne+int(mynumberTwo))

	// //Random Number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
