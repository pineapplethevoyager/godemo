package main

import "fmt"

func main() {
	//Map
	m:= map[string]int{"foo":42}
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)
}