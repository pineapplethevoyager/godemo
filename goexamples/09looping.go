package main

func main() {
	//collection index
	slice := []int{1, 2, 3}   //Could be map
	for i, v := range slice { //or 'i' for keys only or '_, v' for value only
		println(i, v)
	}
}