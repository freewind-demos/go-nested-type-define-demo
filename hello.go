package main

import "fmt"

type User struct {
	name    string
	age     int32
	address [] string
}

func main() {
	user := User{
		name:    "Go",
		age:     10,
		address: []string{"aaa", "bbb"},
	}
	fmt.Println(user)
}
