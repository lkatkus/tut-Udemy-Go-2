package main

import "fmt"

type person struct {
	first, last string
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last)
}

type alien struct {
	first, last, planet string
}

func (a alien) speak() {
	fmt.Println("My name is", a.first, a.last, " and I am from", a.planet)
}

type sentient interface {
	speak()
}

func main() {
	jb := person{
		first: "James",
		last:  "Bond",
	}
	a := alien{
		first:  "Zergus",
		last:   "Maximus",
		planet: "Zerg Prime",
	}

	bar(jb)
	makePersonTalk(jb)

	bar(a)
	makePersonTalk(a)
}

func bar(s sentient) {
	switch s.(type) {
	case person:
		// Because switch is testing for type,
		// then we can assert and use properties from that type
		// s.(person) - type assertion
		fmt.Println("Is type person", s.(person))
	case alien:
		fmt.Println("Is type alien", s.(alien).planet)
	}
}

func makePersonTalk(s sentient) {
	fmt.Println(s)
	s.speak()
}
