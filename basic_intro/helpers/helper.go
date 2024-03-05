package helpers

import (
	"fmt"
	"math/rand"
)

type SomeType struct {
	Typename   string
	Typenumber int
}

func SayHello() {
	fmt.Println("hello")
}

func Random(n int) int {
	val := rand.Intn(n)
	return val

}
