package main

import "fmt"

type user struct {
	name string
	psw  string
}

func (u *user) resetPsw(pas string) {
	u.psw = pas
}
func main() {
	/*
		m := make(map[string]int)
		m["one"] = 1
		m["two"] = 2
		r, ok := m["two"]
		if !ok {
			fmt.Println("Key 'three' do not exist.")
		}
		fmt.Println(r)

		m2 := map[string]int{"one": 1, "two": 2}
		fmt.Println(m2["one"])*/
	//结构体和方法
	a := user{
		name: "xiang",
		psw:  "123",
	}
	a.resetPsw("456")
	fmt.Println(a)
}
