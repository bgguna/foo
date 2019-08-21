package main

import "fmt"

// Create a function called `echo` that takes a single argument
// of type string and prints it out.
// func() {}
func echo(str string) {
	fmt.Println(str)
}

// Change sayHello to allow the greeting to be customized
func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	// Say "Good morning" to John
	echo(sayHello("John"))
}
