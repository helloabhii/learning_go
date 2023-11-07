package main

import "fmt"

func main() {
	defer fmt.Println("Printed after all") // no matter you write this line first or last when you use defer means executes at last
	//it works in LIFO - last in first out
	defer fmt.Println("one")
	defer fmt.Println("two") //after hellobhi - two - one then Printed after all
	fmt.Println("helloabhii")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) /// 0,1,2,3,4 - 4,3,2,1,0
	}
}
