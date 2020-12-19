package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("------ Task 1 ------")
	task1()
	fmt.Println("------ Task 2 ------")
	task2()
	fmt.Println("------ Task 3 ------")
	task3()
	fmt.Println("------ Task 4 ------")
	task4()
	fmt.Println("------ Task 5 ------")
	task5()
	fmt.Println("------ Task 6 ------")
	task6()
	fmt.Println("------ Task 7 ------")
	task7()
	fmt.Println("------ Task 8 ------")
	task8()
	fmt.Println("------ Task 9 ------")
	task9()
	fmt.Println("------ Task 10 ------")
	task10()
}

func task1() {
	a := foo()
	b, c := bar()

	fmt.Println(a)
	fmt.Println(b, c)
}

func foo() int {
	return 99
}

func bar() (int, string) {
	return 1, "problem"
}

func task2() {
	a := []int{1, 2, 3, 4, 5}

	b := otherFoo(a...)
	c := otherBar(a)

	fmt.Println(b, c)
}

func otherFoo(i ...int) int {
	return sum(i)
}

func otherBar(i []int) int {
	return sum(i)
}

func sum(i []int) int {
	sum := 0

	for _, v := range i {
		sum += v
	}

	return sum
}

func task3() {
	defer func() {
		fmt.Println("1")
	}()
	func() {
		fmt.Println("2")
	}()
	func() {
		fmt.Println("3")
	}()
}

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old")
}

func task4() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	m := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   43,
	}

	p.speak()
	m.speak()
}

type square struct {
	side float64
}

func (s square) calculateArea() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) calculateArea() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	calculateArea() float64
}

func task5() {
	s := square{side: 15}
	c := circle{radius: 12.345}

	getArea(s)
	getArea(c)
}

func getArea(s shape) {
	fmt.Printf("%T\t%v\n", s, s.calculateArea())
}

func task6() {
	func() {
		fmt.Println("Another anonymous function!")
	}()
}

func task7() {
	f := func() {
		fmt.Println("I am a function assigned to a variable!")
	}

	f()
}

func task8() {
	f := returner()

	f()
}

func returner() func() {
	return func() {
		fmt.Println("I have been returned by returner function")
	}
}

func task9() {
	nextFoo(func() {
		fmt.Println("I am a callback")
	})
}

func nextFoo(cb func()) {
	fmt.Println("I am nextFoo")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	cb()
}

func task10() {
	a := nextBar()
	b := nextBar()

	fmt.Println("--- A ---")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("--- B ---")
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func nextBar() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
