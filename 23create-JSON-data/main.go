package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Prince   int
	Platform string   `json:"website"`
	Password string   `json:"-"` // - for hiding the password
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

// encoding of JSON
func EncodeJson() {
	lcoCourses := []Course{
		{"Go", 999, "www.google.com", "xyz71", []string{"devOps", "Linux"}},
		{"Jo", 99, "www.google.com", "xy71", []string{"CI/CD", "Linux"}},
		{"Ho", 9, "www.google.com", "x71", nil},
	}

	//package this data in json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // \t for tab space // " " - for space
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
