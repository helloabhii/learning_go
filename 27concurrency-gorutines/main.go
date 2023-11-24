package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointers

func main() {
	// go greeter("hello")
	// greeter("abhii")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://helloabhii.hashnode.dev",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOps in endpoint")
	} else {
		fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
	}
}
