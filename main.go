package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Avatar []string
}

func main() {
	var u = &User{
		Name: "s1ark",
		Age:  20,
		Avatar: []string{
			"https://123.jpg",
			"https://321.jpg",
		},
	}

	byteU, _ := json.Marshal(u)
	fmt.Println(string(byteU))

	var u2 User
	json.Unmarshal(byteU, &u2)
	fmt.Println(u2)
}
