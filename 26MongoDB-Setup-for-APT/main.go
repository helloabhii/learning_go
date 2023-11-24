// go get -u github.com/gorilla/mux
// go get go.mongodb.org/mongo-driver/mongo
// login mongodb /sign in if you haven't
// setup database deployment
// ctrl +shift+p = go tools install and update select all
// at last run
// go build main.go
// go run main.go
// go to thunderclient extension in vs code - new request>type url > http://localhost:4000/api/movie>post> body>json>
//
//	 {
//	 "movie":"Deadpool",
//	  "watched":false
//	}
//
// type url above
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/learning_go/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Serve is getting started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 . . .")
}
