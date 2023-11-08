package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// rt := fmt.Println
const url = "http://lco.dev"

func main() {
	fmt.Println("WEB Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//to check the type of response
	fmt.Printf("Response is of type : %T\n", response)
	defer response.Body.Close() //caller responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
