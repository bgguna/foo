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

	fmt.Printf("%+v\n", user)
	fmt.Println("")

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
}
