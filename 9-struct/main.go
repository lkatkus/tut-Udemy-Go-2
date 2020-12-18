package main

import "fmt"

type person struct {
	name, lastName string
	age            int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	jb := secretAgent{
		person: person{
			name:     "James",
			lastName: "Bond",
			age:      32,
		},
		ltk: true,
	}
	mm := person{
		name:     "Miss",
		lastName: "Moneypenny",
	}

	jb.person.name = "Works"
	fmt.Println(jb.name, jb.person.name)
	jb.name = "Still works"
	fmt.Println(jb.name, jb.person.name)

	fmt.Println(jb, jb.name, jb.person.name, jb.lastName, jb.age, jb.ltk)
	// inner type fields get "promoted" to outter type
	fmt.Println(jb.name, jb.person.name)
	fmt.Println(mm, mm.name, mm.lastName)

	// Anonymous struct
	p1 := struct {
		name, lastName string
		age            int
	}{
		name:     "James",
		lastName: "Bond",
		age:      32,
	}

	fmt.Println(p1)
}
