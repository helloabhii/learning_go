// slices under the hood array
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices !!!")

	var fruits = []string{"S", "T", "U", "V"}
	fmt.Printf("Type of fruits :%T\n ", fruits)

	fmt.Println("Before Append : ", fruits)

	fruits = append(fruits, "W", "X")
	fmt.Println("After append : ", fruits)

	fruits = append(fruits[1:])
	fmt.Println("Exclude 0th index : ", fruits)

	arr2 := make([]int, 3)
	arr2[0] = 71
	arr2[1] = 69
	arr2[2] = 68
	// arr2[3] = 21 // gives error out of bound
	arr2 = append(arr2, 22, 7) //here append realocate the memory
	fmt.Println(arr2)

	sort.Ints(arr2) //sort the slice in ascending order
	fmt.Println("sorted slice : ", arr2)
	fmt.Println(sort.IntsAreSorted(arr2)) //gives true or false

	///////////////////////////////////////////////////////////	///////////////////////////////////////////////////////////
	//remove a slice based on index
	///////////////////////////////////////////////////////////	///////////////////////////////////////////////////////////

	var arr3 = []string{"Python", "Go", "Linux", "Docker", "Kubernetes"}
	fmt.Println(arr3)
	index := 2
	arr3 = append(arr3[:index], arr3[index+1:]...) // means print items till index(2) (2= n-1 =1)//0 & 1 only//
	fmt.Println(arr3)

}
