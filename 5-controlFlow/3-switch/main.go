package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Well it's false")
	case 2 == 2:
		fmt.Println("Obviously...")
		fallthrough
	case false:
		fmt.Println("Still obvious")
		// False, but still falls through
		fallthrough
	case true:
		fmt.Println("But why..?")
	}

	switch {
	case false:
		fmt.Println("Well it's false")
	default:
		fmt.Println("DEFAULT")
	}

	switch {
	case false:
		fallthrough
	default:
		fmt.Println("FUNKY DEFAULT")
	}

	switch "derp" {
	case "notDerp":
		fmt.Println("Nope")
	case "derp":
		fmt.Println("YES!")
	}

	switch "derp" {
	case "notDerp":
		fmt.Println("Nope")
	case "stillNotDerp", "derp":
		fmt.Println("YES AGAIN!")
	}
}
