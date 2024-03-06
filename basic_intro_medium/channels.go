package main

import (
	"fmt"
	"learn/helpers"
)

const numPool = 10

func Cal_val(intChan chan int) {
	random_val := helpers.Random(numPool)
	intChan <- random_val

}

func main() {

	intChan := make(chan int)
	defer close(intChan)

	go Cal_val(intChan)

	num := <-intChan
	fmt.Println(num)
}
