package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang !!!")

	vr := make(map[string]string) // [string] = key, string =  value
	vr["abhi"] = "DevOps"         // key - value
	vr["me"] = "Linux"

	fmt.Println("Map  looks like : ", vr)
	fmt.Println("Abhi equalt to :", vr["abhi"]) // when we pass key here it result its value

	// to remove the any item
	delete(vr, "abhi")
	fmt.Println(vr)

	//loops
	for key, value := range vr {
		fmt.Printf("For key  %v, value is %v\n", key, value)
	}
	for _, value := range vr {
		fmt.Printf("For key  _, value is %v\n", value)
	}
}
