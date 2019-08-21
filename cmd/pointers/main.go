package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)
}

func printValue(s *string) {
	// print the pointer value
	fmt.Printf("s*: %v\n", s)
	// print the string value
	fmt.Printf("s: %s\n", *s)
	// change the string value
	*s = "Ringo"
}
