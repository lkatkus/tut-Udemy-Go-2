package main

import "fmt"

func main() {
	fmt.Println("------ Task 1 ------")
	task1()
	fmt.Println("------ Task 2 ------")
	task2()
}

func task1() {
	x := 999

	fmt.Println(x, &x)
}

type person struct {
	first, last string
	age         int
}

func (p *person) changeMe() {
	p.first = "Zerg"
	(*p).last = "Other"
	p.age = 99
}

func (p person) otherChange() {
	p.first = "Next Name"
	p.age = 123

	fmt.Println("Will not work", p)
}

func task2() {
	p := person{
		first: "Person",
		last:  "Namer",
		age:   34,
	}

	fmt.Println("Before", p)

	p.changeMe()
	fmt.Println("After", p)

	p.otherChange()
	fmt.Println("After not working", p)
}
