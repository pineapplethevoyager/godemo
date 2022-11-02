package main

//GO only has for loops
func main() {
	println("Loop");
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}

	println("Loop with clauses");
	for c := 0; c < 5; c++ {
		println(c)
	}

	println("Infinite Loop");
	var x int
	for { // or for ;;{}
		if x == 5 {
			break
		}
		println(x)
		x++
	}
}