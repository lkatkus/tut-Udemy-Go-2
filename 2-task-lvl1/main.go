package main

import "fmt"

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
}

func task1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("x is equal to %v and is of type %T\n", x, x)
	fmt.Printf("y is equal to %v and is of type %T\n", y, y)
	fmt.Printf("z is equal to %v and is of type %T\n", z, z)
}

func task2() {
	var x int
	var y string
	var z bool

	// Printing zero values
	fmt.Printf("x is equal to %v and is of type %T\n", x, x)
	fmt.Printf("y is equal to %v and is of type %T\n", y, y)
	fmt.Printf("z is equal to %v and is of type %T\n", z, z)
}

func task3() {
	var x int = 42
	var y string = "James Bond"
	var z bool = true

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

func task4() {
	type myType int

	var x myType
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v\t%T\n", x, x)
}

func task5() {
	type myType int
	var x myType
	var y int

	y = int(x)

	fmt.Printf("y is equal to %v and of type %T\n", y, y)
}
