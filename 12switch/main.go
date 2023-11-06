package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Golang ")

	rand.Seed(time.Now().UnixNano()) // to get the random values
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is : ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("You Got 1 point")
	case 2:
		fmt.Println("You Got 2 points")
	case 3:
		fmt.Println("You Got 3 points")
		fallthrough //means next case also executes
	case 4:
		fmt.Println("You Got 4 point")
	case 5:
		fmt.Println("You Got 5 points")
	case 6:
		fmt.Println("You Got 6 points")
	default:
		fmt.Println("What was that !!!")
	}

}
