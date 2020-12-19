package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Age         int    `json:"age"`
	notExported string
}

func main() {
	people := []person{
		person{
			Name:        "First name",
			LastName:    "First lastName",
			Age:         11,
			notExported: "Will not export, because is lowercase",
		},
		person{
			Name:        "Second name",
			LastName:    "Second lastName",
			Age:         22,
			notExported: "Will not export, because is lowercase",
		},
	}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--- struct ---")
	fmt.Println(people)
	fmt.Println("--- JSON ---")
	// json.Marshal returns []byte, which needs to be converted to string
	convertedBs := string(bs)
	fmt.Println(convertedBs)

	getStructFromJSON()
}

type importPerson struct {
	TargetName     string `json:"name"`
	TargetLastName string `json:"lastName"`
	TargetAge      int    `json:"age"`
}

func getStructFromJSON() {
	jsonString := `[{"Name":"First name","LastName":"First lastName","Age":11},{"Name":"Second name","LastName":"Second lastName","Age":22},{"Name":"Third name","LastName":"Third lastName","Age":99}]`

	var peopleFromJSON []importPerson

	fmt.Println("--------- BEFORE ---------")
	fmt.Println(peopleFromJSON)
	json.Unmarshal([]byte(jsonString), &peopleFromJSON)

	fmt.Println("--------- AFTER ---------")
	fmt.Printf("%+v\n", peopleFromJSON)

	for _, p := range peopleFromJSON {
		fmt.Printf("%+v\n", p)
	}
}
