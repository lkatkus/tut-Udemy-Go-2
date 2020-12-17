package main

import (
	"fmt"
	"time"
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
	task6(false)
	task6(true)
	fmt.Println("------ Task 7 ------")
	task7(false)
	task7(true)
	fmt.Println("------ Task 8 ------")
	task8()
	fmt.Println("------ Task 9 ------")
	task9()
}

func task1() {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Println()
}

func task2() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)

		for j := 65; j <= 90; j++ {
			fmt.Printf("\t%U\t%c\n", j, j)
		}
	}
}

func task3() {
	bd := 1987
	currentYear := time.Now().Year()

	for i := bd; i <= currentYear; i++ {
		fmt.Println(i)
	}
}

func task4() {
	bd := 1987
	currentYear := time.Now().Year()

	for {
		if bd > currentYear {
			break
		}

		fmt.Println(bd)
		bd++
	}
}

func task5() {
	for i := 10; i <= 23; i++ {
		m := fmt.Sprint(i % 4)

		fmt.Printf("When %v is divided by %v the modulus is %v\n", i, 4, m)
	}
}

func task6(b bool) {
	if b {
		fmt.Println("Yes it's true")
	}
}

func task7(b bool) {
	if b {
		fmt.Println("Yes it's true")
	} else {
		fmt.Println("Try again")
	}
}

func task8() {
	switch {
	case false:
		fmt.Println("FALSE")
	case true:
		fmt.Println("TRUE")
	}
}

func task9() {
	switch "favSport" {
	case "notAnotherSwitch":
		fmt.Println("FALSE")
	case "favSport":
		fmt.Println("YES IT IS TRUE")
	}
}
