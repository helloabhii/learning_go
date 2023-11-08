package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Handling URL in golang")
	fmt.Println(myurl)
	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of the qparams is : %T\n ", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is : ", val)
	}
	// for key, val := range qparams {
	// 	fmt.Println("Params is : ", key, val)
	// }
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user= helloabhii",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
