package main

import "fmt"

// Everything that is out of a function in golang, must be
// declared with a keyword, for example 'var'
// a := 2 -> that, for example it's impossible to do, Golang not accept declaration and set at same time out of function
// var x -> but, creating a variable and setting the value in a function is not a problem

func process() {
	//multiple assingment example
	a, b := 10, "foo"

	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	process()
}
