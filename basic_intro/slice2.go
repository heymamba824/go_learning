package main

import "fmt"

type myStruct struct {
	FirstName string
}

func call(m *myStruct) string {
	return m.FirstName
}
func (m *myStruct) struct2() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Xiang"
	another := myStruct{
		FirstName: "Zhu",
	}
	fmt.Println("my var is set to", myVar.FirstName)
	fmt.Println("my var is set to", another.FirstName)
	fmt.Println("my var is set to", call(&myVar))
	fmt.Println("my var is set to", call(&another))
	fmt.Println("my var is set to", myVar.struct2())
}
