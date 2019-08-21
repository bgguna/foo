package main

import (
	"fmt"

	"github.com/bgguna/foo"
)

func main() {
	user := foo.User{
		First: "Bogdan",
		Last:  "Guna",
	}

	fmt.Println("1. Packages")
	fmt.Printf("%+v\n", user)
	fmt.Println("")

	// Arrays
	fmt.Println("2. Arrays")
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawbery"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%d %s\n", i, fruits[i])
	}

	fmt.Println("")

	for i, fruit := range fruits {
		fmt.Printf("%d %s\n", i, fruit)
	}

	fmt.Println("")

	for i, fruit := range fruits {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d %s\n", i, fruit)
	}

	fmt.Println("")

	// Slices
	fmt.Println("3. Slices")
	fruitsSlice := []string{"Banana", "Orange", "Pineapple", "Strawbery"}
	for i, fruit := range fruitsSlice {
		fmt.Printf("%d %s\n", i, fruit)
	}
	fmt.Println("")

	// The make keyboard allows us to define the starting length of the slice, and optionally, the starting capacity of the slice.
	a := make([]int, 1, 3)
	fmt.Println(a)      // [0]
	fmt.Println(len(a)) // 1
	fmt.Println(cap(a)) // 3

	fmt.Println("")
	// make and append
	b := make([]string, 2)
	b = append(b, "foo", "bar")
	fmt.Printf("%q\n", b) // ["" "" "foo" "bar"]

	fmt.Println("")
}
