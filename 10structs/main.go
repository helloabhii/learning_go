package main

import "fmt"

func main() {
	fmt.Println("Struct in golang !!!")
	//no inheritence in golang
	// no super or parent class

	abhi := User{"Abhishek", "ak8888780@gmail.com", true, 21}
	fmt.Println(abhi)
	fmt.Printf("abhi details are: %+v\n", abhi)
	fmt.Printf("Name is %v and  email is  %v.", abhi.Name, abhi.Email) //

}

type User struct { // defined here
	Name   string
	Email  string
	Status bool
	Age    int
}
