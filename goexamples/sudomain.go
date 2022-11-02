package main

import "fmt"
const (
	first = iota + 6
	second = 2 << iota //bit shift 2 by 1 
)
const (
	zero = iota
	one
)

func main() {
	//Variable
	var i int
	i = 42
	fmt.Println(i)

	chicken := "food or friend?"
	fmt.Println(chicken)

	//Pointer
	var firstname *string = new(string)
	*firstname = "chris"
	fmt.Println(*firstname);

	var deref string = "bob";
	ptr := &deref 
	*ptr = "joe"
	fmt.Println(ptr, *ptr)
	//Const
	const num = 3 // must be at compile time
	//Implicit type vs defined type
	fmt.Println(num + 3)
	fmt.Println(num + 1.2)

	//array
	var arr [3]int
	arr[0] = 1
	fmt.Println(arr)

	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)

	//slice
	slice0 := arr2[:]
	slice0 = append(slice0, 4,5,6)//changes array and slice

	slice := []int{1,2,3}
	slice = append(slice, 4,5,6)
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2,s3,s4)

	//Map
	m:= map[string]int{"foo":42}
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)

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