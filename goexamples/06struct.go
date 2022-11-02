package main

import "fmt"

func main() {
	//struct
	type user struct { // can be moved to any scope
		ID int
		FirstName string
		LastName string
	}
	var u user
	u.ID = 1
	u.FirstName = "Chris"
	u.LastName = "Lopez"
	u2:= user{ ID: 1,
		FirstName: "Chris",
		LastName: "lopez",
	}
	fmt.Println(u2)
}