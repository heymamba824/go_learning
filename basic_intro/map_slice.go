// 可以存任何东西去map
package main

import (
	"fmt"
	"sort"
)

type User struct {
	firstName string
	lastName  string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Jap"
	fmt.Println(myMap["cat"])
	fmt.Println(myMap["dog"])

	game := make(map[string]User)
	pass := User{
		firstName: "mike",
		lastName:  "dada",
	}
	game["xiang"] = pass
	fmt.Println(game["xiang"])

	var mySlice []string
	mySlice = append(mySlice, "game 1 ")
	mySlice = append(mySlice, "game2")
	fmt.Println(mySlice)
	fmt.Println(mySlice[1])
	numbers := []int{1, 32, 314, 1231, 1, 3, 141, 31, 3}
	sort.Ints(numbers)
	fmt.Println(numbers)
}
