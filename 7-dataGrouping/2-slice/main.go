package main

import "fmt"

func main() {
	// composite literal - type{values}
	// type - []int
	// values : {1, 2, 3, 4, 5}

	x := []int{1, 2, 3, 42, 5}

	fmt.Println("x: ", x)
	fmt.Println("x[2]: ", x[2])
	fmt.Println("len(x): ", len(x))

	fmt.Println("x[2:]: ", x[2:])   // from including to the end
	fmt.Println("x[:2]: ", x[:2])   // up to and not include
	fmt.Println("x[1:3]: ", x[1:3]) // from including to not including

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	for i, v := range x {
		fmt.Println(i, v)
	}

	x = append(x, 99, 111, 112)
	fmt.Println(x)

	y := []int{67, 44}

	x = append(x, y...)
	fmt.Println(x)

	x = append(x, []int{999, 991}...)
	fmt.Println(x)

	// deleting from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// type, length, capacity
	z := make([]int, len(x), 11)

	// when appending to slice which still has capacity a new array is not created
	// this is saving processing time
	z = append(z, 1234)
	// when appending outside of capacity a new array is created
	z = append(z, 1234222)

	fmt.Println(len(z))
	fmt.Println(cap(z))

	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}

	xp := [][]string{jb, mp}

	fmt.Println(jb)
	fmt.Println(mp)
	fmt.Println(xp)
}
