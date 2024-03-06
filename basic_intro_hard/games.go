package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("the random number is", secretNumber)
	fmt.Println("game starts, input a number between 0-100")
	for {
		reader := bufio.NewReader(os.Stdin)
		//读取一行输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			//return
			continue
		}
		//去掉换行符
		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("your guess is", guess)
		if guess == secretNumber {
			fmt.Println("congradualations")
			return
		} else if guess > secretNumber {
			fmt.Println("The number is less than your guess")
		} else {
			fmt.Println("The number is bigger than your guess")
		}

	}

}
