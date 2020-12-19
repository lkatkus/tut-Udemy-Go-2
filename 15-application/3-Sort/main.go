package main

import (
	"fmt"
	"sort"
)

func main() {
	basicSort()
	customSort()
}

func basicSort() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("Before: ", xi)
	sort.Ints(xi)
	fmt.Println("After: ", xi)

	fmt.Println("Before: ", xs)
	sort.Strings(xs)
	fmt.Println("After: ", xs)
}

type person struct {
	first string
	age   int
}

type byAge []person

// to implement sort package interface
func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

type byName []person

// to implement sort package interface
func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].first < a[j].first }

func customSort() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}
	p5 := person{"Adam", 56}

	people := []person{p1, p2, p3, p4, p5}

	fmt.Println("BEFORE byAge", people)
	sort.Sort(byAge(people))
	fmt.Println("AFTER", people)

	fmt.Println("BEFORE byName", people)
	sort.Sort(byName(people))
	fmt.Println("AFTER", people)
}
