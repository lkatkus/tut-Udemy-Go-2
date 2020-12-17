package main

import (
	"fmt"
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
}

func task1() {
	for i := 15; i < 24; i++ {
		fmt.Println(getFormattedString(i))
	}
}

func getFormattedString(i int) string {
	return fmt.Sprintf("%v\t%b\t%x", i, i, i)
}

func task2() {
	a := 42 == 42
	b := 42 <= 53
	c := 42 >= 53
	d := 42 != 53
	e := 42 < 53
	f := 42 > 53

	fmt.Println("42 == 42\t", a)
	fmt.Println("42 <= 53\t", b)
	fmt.Println("42 >= 53\t", c)
	fmt.Println("42 != 53\t", d)
	fmt.Println("42 < 53\t\t", e)
	fmt.Println("42 > 53\t\t", f)

}

func task3() {
	const (
		// Untyped constant
		a = 42
		// Typed constant
		b int = 32
	)

	fmt.Println(a, b)
}

func task4() {
	a := 42
	fmt.Printf("%v\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%v\t%b\t%#x\n", b, b, b)
}

func task5() {
	rawString := `
	Very raw
		very string
			very indented
`

	fmt.Println(rawString)
}

func task6() {
	const start = 2000

	const (
		a = start + iota
		b = start + iota
		c = start + iota
		d = start + iota
	)

	fmt.Println(a, b, c, d)
}
