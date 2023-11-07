package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	abhi := User{"Abhishek", "ak8888780@gmail.com", true, 21}
	fmt.Println(abhi)
	abhi.GetStatus()
	abhi.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active : ", u.Status)

}
func (u User) NewMail() { // struct passes on the copies
	u.Email = "test@123go.dev"
	fmt.Println("Email of this user is : ", u.Email)
}
