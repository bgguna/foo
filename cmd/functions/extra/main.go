package main

import "fmt"

// Create a "compare" function that takes two integers and compares them using
// the following function definition:
//
// func less(i, j int) bool {
//    return i < j
// }
// print out if i is less than j
// func () {}

type less func(i, j int) bool

func compare(a, b int, l less) {
	fmt.Printf("%v is less  than %v: %v\n", a, b, l(a, b))
}

func main() {
	// Call compare function with number pairs "0, 0", "0, 1", "1, 0"
	ls := func(i, j int) bool {
		return i < j
	}

	compare(0, 0, ls)
	compare(0, 1, ls)
	compare(1, 0, ls)
}
