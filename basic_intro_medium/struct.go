package main

import (
	"fmt"
	"time"
)

type Person1 struct {
	firstName string
	lastName  string
	phone     string
	age       int
	birth     time.Time
}

func main() {
	user := Person1{
		firstName: "xiang",
		lastName:  "zhu",
	}
	fmt.Println(user.firstName, user.lastName)
}
