package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Last    string `json:"last"`
	Hair    string `json:"hair"`
	Has_dog bool   `json:"has_dog"`
}

func main() {
	// write json to a struct
	myJson := `
[
	{
		"name" : "xiang",
		"last" : "zhu",
		"hair" : "black",
		"has_dog" : true
	},
	{
		"name" : "xq",
		"last" : "zhu",
		"hair" : "yellow",
		"has_dog": false
	}
]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
	}
	fmt.Printf("unmarshalled: %v\n", unmarshalled)
	// write json from a struct
	var mySlice []Person
	var marsh1 Person
	marsh1.Name = "hero"
	marsh1.Hair = "white"
	marsh1.Has_dog = false
	marsh1.Last = "super"

	mySlice = append(mySlice, marsh1)

	var marsh2 Person
	marsh2.Name = "lucas"
	marsh2.Hair = "077"
	marsh2.Has_dog = false
	marsh2.Last = "talent"

	mySlice = append(mySlice, marsh2)
	//第二个参数是每一行 JSON 的前缀字符串，在这个例子中是 ""（空字符串），表示没有前缀。
	//第三个参数是每一层 JSON 结构的缩进字符串，在这个例子中也是 ""（空字符串），表示不添加缩进。如果你想要格式化输出，可以提供如 " "（一个空格）或 "\t"（一个制表符）作为缩进字符串。
	newJson, err := json.MarshalIndent(mySlice, "", "\t")
	if err != nil {
		fmt.Println("Error unmarshalling", err)
	}
	fmt.Printf(string(newJson))
	//fmt.Print(newJson)
}
