package main

import "fmt"

type Animal interface {
	Says() string
	Numberoflegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name   string
	Color  string
	Number int
}

func Printinfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "and has", a.Numberoflegs())

}
func (d *Dog) Numberoflegs() int {
	return 4
}

func (d *Dog) Says() string {
	return d.Breed
}
func (g *Gorilla) Numberoflegs() int {
	return 88
}

func (g *Gorilla) Says() string {
	return "heyyyy"
}
func main() {
	dog := Dog{
		Name:  "kakakakak",
		Breed: "German",
	}
	Printinfo(&dog)

	gorilla := Gorilla{
		Name:   "Joker",
		Color:  "white",
		Number: 100,
	}
	Printinfo(&gorilla)
}
