package main

import "fmt"

func main() {
	fmt.Println(mySum(1, 2, 3))
	fmt.Println(mySum(4, 4))
	fmt.Println(mySum(5, 95))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
