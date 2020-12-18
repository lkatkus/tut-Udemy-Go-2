package main

import "fmt"

func main() {
	// type - map[string]int, where string is key, int is value
	m := map[string]int{
		"James":           32,
		"Miss MoneyPenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	// If key is missing, then it returns a ZERO VALUE
	fmt.Println(m["Missing"])
	// second return argument is false if key exists
	v, ok := m["Missing"]
	fmt.Println(v, ok)

	if v, ok := m["OtherMissingKey"]; ok {
		fmt.Println("EXISTING", v)
	} else {
		fmt.Println("MISSING")
	}

	if v, ok := m["Miss MoneyPenny"]; ok {
		fmt.Println("EXISTING", v)
	} else {
		fmt.Println("MISSING")
	}

	// adding new item
	m["Zergas"] = 9999
	fmt.Println(m)

	// print key and value
	for k, v := range m {
		fmt.Println(k, v)
	}

	// print only keys
	for k := range m {
		fmt.Println(k)
	}

	fmt.Println(m)
	delete(m, "Zergas")
	fmt.Println(m)
}
