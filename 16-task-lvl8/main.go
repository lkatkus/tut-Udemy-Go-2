package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

var u1 = user{
	First: "James",
	Last:  "Bond",
	Age:   32,
	Sayings: []string{
		"Shaken, not stirred",
		"Youth is no guarantee of innovation",
		"In his majesty's royal service",
	},
}

var u2 = user{
	First: "Miss",
	Last:  "Moneypenny",
	Age:   27,
	Sayings: []string{
		"James, it is soo good to see you",
		"Would you like me to take care of that for you, James?",
		"I would really prefer to be a secret agent myself.",
	},
}

var u3 = user{
	First: "M",
	Last:  "Hmmmm",
	Age:   54,
	Sayings: []string{
		"Oh, James. You didn't.",
		"Dear God, what has James done now?",
		"Can someone please tell me where James Bond is?",
	},
}

func task1() {
	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(bs))
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func task2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var people []person

	fmt.Println("--- BEFORE ---")
	fmt.Println(people)
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("--- AFTER ---")
	for _, p := range people {
		fmt.Println(p)
	}
}

func task3() {
	users := []user{u1, u2, u3}

	fmt.Println(users)

	ec := json.NewEncoder(os.Stdout)
	err := ec.Encode(users)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func task4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type byAge []user

// to implement sort package interface
func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []user

// to implement sort package interface
func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].First < a[j].First }

func task5() {
	users := []user{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(byAge(users))
	fmt.Println(users)

	fmt.Println(users)
	sort.Sort(byName(users))
	fmt.Println(users)
}
