package main

import "fmt"

//只有for loop

func main1() {
	//for i := 0; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	// enumrate
	animals := []string{"dog", "fish", "horse", "cow"}
	for i, animal := range animals {
		fmt.Println(i, animal)
	}
	for animal := range animals {
		fmt.Println(animal)
	}
	for _, animal := range animals {
		fmt.Println(animal)
	}
}

func main() {
	animals := make(map[string]string)
	animals["dogs"] = "Fido"
	animals["cat"] = "Kat"
	for animal, name := range animals {
		fmt.Println(animal, name)
	}
}
