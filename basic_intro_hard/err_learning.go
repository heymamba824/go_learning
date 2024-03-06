package main

import (
	"errors"
	"fmt"
)

type user1 struct {
	name     string
	password string
}

func FindUser(users []user1, name string) (*user1, error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("Not find the user name")
}
func main() {
	u := []user1{
		{"Alice", "alice123"},
		{"Bob", "bob123"},
	}
	u = append(u, user1{"xiang", "123"})
	exist, err := FindUser(u, "xiang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exist.name)
	if u1, err := FindUser(u, "wang"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u1.name)
	}
}
