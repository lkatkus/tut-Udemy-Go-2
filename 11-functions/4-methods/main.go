package main

import "fmt"

type person struct {
	first, last string
}

func (p person) sayName() {
	fmt.Println("My name is", p.first, p.last)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) doStuff() {
	fmt.Println(sa.first, "doing secret agent stuff")
}

func main() {
	jb := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	mm := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	jb.sayName()
	jb.doStuff()

	mm.sayName()
}

func foo() {
	fmt.Println("foo")
}
