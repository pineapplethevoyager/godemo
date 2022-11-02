package main

import "fmt"

func main() {
	var firstname *string = new(string)
	*firstname = "chris"
	fmt.Println(*firstname);

	var deref string = "bob";
	ptr := &deref 
	*ptr = "joe"
	fmt.Println(ptr, *ptr)
}