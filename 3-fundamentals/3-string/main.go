package main

import "fmt"

func main() {
	const SomeConst = 1
	const (
		a int     = 12
		b float64 = 13.99
		c string  = "99"
		// Untyped constant
		d = 123
	)

	s := "Some string"
	const MyConst = 123

	fmt.Printf("%v\t%T\n", s, s)

	// string is a slice of byte
	bs := []byte(s)
	fmt.Printf("%v\t%T\n", bs, bs)
	fmt.Println(s[0], string(s[0]))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for _, v := range s {
		// string, binary, hex, utf8
		fmt.Printf("%v\t%v\t%#x\t%#U\n", string(v), v, v, v)
	}
}
