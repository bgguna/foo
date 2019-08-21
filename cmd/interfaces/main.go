package main

import "fmt"

type User struct {
	ID    int
	First string
	Last  string
}

// example:
//      User 1 is Rob Pike
func (u User) String() string {
	return fmt.Sprintf("User %d is %s %s", u.ID, u.First, u.Last)
}

func main() {
	u := User{ID: 1, First: "Rob", Last: "Pike"}
	fmt.Println(u)
}
