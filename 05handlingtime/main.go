package main

import (
	"fmt"
	"time" //have to import time to use it
)

func main() {
	fmt.Println("Here we working with built-in time function")
	presentTime := time.Now() //putting time func into variable called presentTime
	fmt.Println("Current Time : ", presentTime)

	//change the format of the time
	fmt.Println("Current Date : ", presentTime.Format("01-02-2006  15:04:05 Monday")) // this is defined in the documentation to use only this time to get accurate result
}
