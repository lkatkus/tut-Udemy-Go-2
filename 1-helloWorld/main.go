package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// variables declared package level scope
var outsideName = "Outsider"

// DECLARE "z" of type int
// ASSING the ZERO VALUE of type int to z. In this case - 0
var z int

type hotdog int

var intB int = 13
var b hotdog

func main() {
	// short declation operator
	name := "Anonymous"
	fmt.Println("Hello again...", name, outsideName)

	func() {
		fmt.Println("Anonymous function call")
	}()

	b = 99
	fmt.Printf("Custom  typed b - %T, %v \n", b, b)
	b = b + hotdog(intB)
	fmt.Printf("Hotdog + converted int - %T, %v \n", b, b)

	foo()

	for i := 0; i < 5; i++ {
		fmt.Print("i is ", i, "\t")

		s := fmt.Sprintf("%T\t%b\t%x", i, i, i)
		fmt.Println(s)
	}

	bar()
}

func foo() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Say something: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if len(text) == 0 {
		fmt.Println("Don't be shy")
	} else {
		fmt.Printf("You said \"%s\"\n", text)
	}

}

func bar() {
	fmt.Println("and we exit with outsideName", outsideName, z)
}
