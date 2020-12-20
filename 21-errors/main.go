package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	log.SetOutput(f)

	foo()
	bar()
}

func foo() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("James Bond")
	io.Copy(f, r)

	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

func bar() {
	_, err := os.Open("missing-file.txt")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
}
