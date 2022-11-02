package main

import "fmt"

func main() {
	var firstname *string = new(string)
	*firstname = "chris"
	fmt.Println(*firstname);
}