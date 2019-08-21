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
}
