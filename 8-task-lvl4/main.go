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
	fmt.Println("------ Task 6 ------")
	task6()
	fmt.Println("------ Task 7 ------")
	task7()
	fmt.Println("------ Task 8 ------")
	m := task8()
	fmt.Println("------ Task 9 ------")
	task9(m)
	fmt.Println("------ Task 10 ------")
	task10(m)

	fmt.Println("------ AFTER ------")
	fmt.Println(m)
}

func task1() {
	a := [5]int{11, 22, 33, 44, 55}

	fmt.Printf("%T\n", a)

	for i, v := range a {
		fmt.Println(i, v)
	}
}

func task2() {
	a := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	fmt.Printf("%T\n", a)

	for i, v := range a {
		fmt.Println(i, v)
	}
}

func task3() {
	a := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	a1 := a[:5]
	a2 := a[5:]
	a3 := a[2:8]

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

func task4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, []int{53, 54, 55}...)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func task5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func task6() {
	x := make([]string, 50, 50)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func task7() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooo, James"}}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i)
		for _, j := range v {
			fmt.Printf("\t%v\n", j)
		}
	}
}

func task8() map[string][]string {
	m := map[string][]string{}
	fmt.Println(m)

	m["bond_james"] = []string{"Shaken, not stirred", "Martini", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println("Record for:", k)

		for i, t := range v {
			fmt.Println(i, t)
		}
	}

	return m
}

func task9(m map[string][]string) {
	fmt.Println(m)

	m["fleming_ian"] = []string{"Stakes", "Cigars", "Espionage"}
	fmt.Println(m)
}

func task10(m map[string][]string) {
	fmt.Println(m)

	if _, ok := m["bond_james"]; ok {
		fmt.Println("bond_james is present")
	}

	delete(m, "bond_james")

	if _, ok := m["bond_james"]; ok {
		fmt.Println("bond_james is present")
	} else {
		fmt.Println("no more bond_james!")
	}
}
