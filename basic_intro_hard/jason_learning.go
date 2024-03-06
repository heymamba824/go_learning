package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(buf)
	fmt.Println(string(buf))
	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(buf)
	fmt.Println(string(buf))
	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", b)

}
