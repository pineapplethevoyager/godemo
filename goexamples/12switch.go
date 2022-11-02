package main

func main() {
	name := "Michael"
	switch name { //switch has default breaking so no need for it
	case "Carissa":
		println("name is Carissa")
	case "Chris":
		println("name is Chris")
	case "Michael":
		println("name is Michael")
	default:
		println("name not known")
	}
}