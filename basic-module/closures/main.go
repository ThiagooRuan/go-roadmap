package main

import (
	"fmt"
)

// Closures are functions that capture references to
// the variables in the scope where they were defined
// they can access and modify these variables
// even if that scope goes out of execution.
func makeIdGenerator() func() int {
	id := 0 // that is, every time I call the function, it will change and keep the value of the id
	return func () int {
		id++
		return id
	}
}

func main() {
	generateId1 := makeIdGenerator()
	generateId2 := makeIdGenerator()

	fmt.Println(
		generateId1(),
		generateId1(),
		generateId1(),
	)

	// if the callable function is not the same, the values then will be 
	// return aren't be the same as well
	fmt.Println(
		generateId1(),
		generateId2(),
	)
}