package main

import (
	"errors"
	"fmt"
)

func main() {
	a,b:=startWebServer(8080)
	fmt.Println(a)
	fmt.Println(b.Error())
}

func startWebServer(port int) (bool, error) {
	return true, errors.New("Error not implemented")
}