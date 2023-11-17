package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" Golang mod/module")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", r))

}

//before excute run - go mod init github.com/helloabhii/learning_go/24mod // by this go.mod file created

// then run following command - go get -u github.com/gorilla/mux // by this go.sum file created
func greeter() {
	fmt.Println("Hey there mod users")
}
func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to golang helloabhii</h1>"))
}

//go build .
//go mod verify - to check verification of files modules
// go list
// go list all
// go list -m all
// go list  -m -versions  github.com/gorilla/mux
// go mod why github.com/gorilla/mux
// go mod graph
//go mod edit -go 1.8
