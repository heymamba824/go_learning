package main

import (
	"fmt"
	"learn/helpers"
)

func main() {
	fmt.Println("hello")
	var myVar helpers.SomeType
	myVar.Typename = "xiang"
	fmt.Println(myVar.Typename)
	myVar.Typenumber = 99
	fmt.Println("number", myVar.Typenumber)
	helpers.SayHello()
}
