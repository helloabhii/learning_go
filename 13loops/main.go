package main

import "fmt"

func main() {
	fmt.Println("Loops in golang ")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)
	//////////////////////////////////////////////////////////////////////////////////////////
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	//////////////////////////////////////////////////////////////////////////////////////////
	for i := range days {
		fmt.Println(days[i])
	}
	//////////////////////////////////////////////////////////////////////////////////////////
	for index, day := range days { //
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
	// ////////////////////////////////////////////////////////////////////////////////////////
	val := 1
	for val < 11 {
		// if val == 9 {
		// 	break
		// }
		if val == 7 {
			goto gto
			val++
			continue //skip the 7

		}
		fmt.Println(val)
		val++
	}
gto:
	fmt.Println("Goto to Used Here !!!")
}
