package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

func task1() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}

	fmt.Println(string(bs))
}

func task2() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to marshal in toJSON: %v", err)
	}

	return bs, err
}

type customErr struct {
	info string
}

func (e customErr) Error() string {
	return fmt.Sprintf("This is my custom error: %v", e.info)
}

func task3() {
	e := customErr{
		info: "yes yes much error",
	}

	foo(e)
}

func foo(e error) {
	fmt.Println(e.(customErr).info) // type assertion
	fmt.Println("foo ran with: ", e)
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func task4() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			lat:  "50 N",
			long: "99 W",
			err:  errors.New("Quick maths"),
		}
	}
	return 42, nil
}
