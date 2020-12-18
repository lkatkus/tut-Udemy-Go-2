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
}

func task1() {
	type person struct {
		name, lastName string
		favorites      []string
	}

	p1 := person{
		name:      "Namer",
		lastName:  "Lastoner",
		favorites: []string{"Vanila", "Strawberry"},
	}
	p2 := person{
		name:      "Seconder",
		lastName:  "Personer",
		favorites: []string{"Chocolate", "Kiwi"},
	}

	fmt.Println(p1.name)
	fmt.Println(p1.lastName)
	for k, v := range p1.favorites {
		fmt.Println(k, v)
	}

	fmt.Println(p2.name)
	fmt.Println(p2.lastName)
	for k, v := range p2.favorites {
		fmt.Println(k, v)
	}
}

func task2() {
	type person struct {
		name, lastName string
		favorites      []string
	}

	p1 := person{
		name:      "Namer",
		lastName:  "Lastoner",
		favorites: []string{"Vanila", "Strawberry"},
	}
	p2 := person{
		name:      "Seconder",
		lastName:  "Personer",
		favorites: []string{"Chocolate", "Kiwi"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.name)
		fmt.Println(v.lastName)

		for _, f := range v.favorites {
			fmt.Printf("\t%v", f)
		}

	}
}

func task3() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: false,
	}

	fmt.Printf("%T\t%v\n", t, t.doors)
	fmt.Printf("%T\t%v\n", s, s.doors)
}

func task4() {
	type person struct {
		name, lastName string
		age            int
		favorites      []string
		relatives      map[string]person
	}

	p := struct {
		name, lastName string
		age            int
		favorites      []string
		relatives      map[string]person
	}{
		name:     "Zerg",
		lastName: "Ling",
		age:      1,
		favorites: []string{
			"Humans",
			"Marines",
			"Protons",
		},
		relatives: map[string]person{
			"Zergling": person{
				name:     "Zergling",
				lastName: "Zerginator",
				age:      999,
			}},
	}

	fmt.Println(p)
}
