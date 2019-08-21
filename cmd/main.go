package main

import "fmt"

func main() {
	// Maps
	beatles := map[string]string{
		"John":   "guitar",
		"Paul":   "bass",
		"George": "guitar",
		"Ringo":  "drums",
	}

	fmt.Println(beatles)
	fmt.Println()

	for key, value := range beatles {
		fmt.Printf("%v \tplays the %v\n", key, value)
	}
	fmt.Println()

	delete(beatles, "George")
	fmt.Println(beatles)
	fmt.Println()

	paul := beatles["Paul"]
	fmt.Println(paul) // bass
	fmt.Println()

	key := "Paul"
	if value, ok := beatles[key]; ok {
		fmt.Printf("Found key %q: %q", key, value) // Found Key "Paul": "bass"
	} else {
		fmt.Printf("Key not found: %q", key) // Key not found: "foo"
	} // Change the "key" to "foo" and re-run the example

	fmt.Println()
}
