package main

import "fmt"

func process() {
	// here if i don't comment one of this var declaration,
	// the statement in 13 line in invalid
	// because i only can use ':=' when i do var declaration e set the value at same time
	// unless i'm creating a new variable, then it will be accepted

	var (
		a int
		b string
	)

	a, b, c := 10, "foo", -1

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	process()
}
