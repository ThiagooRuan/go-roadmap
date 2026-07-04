package main

import "fmt"

func main() {
	// arrays always has a specific lenght
	// index 0 until len(arr)-1
	// arr = [1, 2, 3]
	//		  0, 1, 2
	// ways to declare arrays
	// var arr [3]int
	// var arr2 = [3]int{1, 2, 3}
	// arr3 := [3]int{}

	var (
		x = [3]int{0, 1, 2}
		y = [4]int{0, 1, 2, 4}
	)

	fmt.Println(x)
	fmt.Println(y)

	// arrays lenght also it's part of the array type
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	// [3]int
	// [4]int

	// setting 5 at index 2
	// x = [1, 2, 3]
	x[2] = 5
	// x = [1, 2, 5]
	fmt.Println(x)

	//
	var arrayCopy [3]int
	arrayCopy = x

	// passing by value or reference.
	// it is possible to copy data from one array to another, but depending on the size of the array
	// it can greatly impact the application's resource cost
	// unlike slices which, being pointer-based, are "passed" by reference,
	// making them less costly for the application
	// this even extends to function calls where you pass an array to the function and modify it, if you don't return that value,
	// the array will only be changed inside the function, but outside, it will keep its previous value. Using slice is different,
	// because the function receives and changes the reference, so even without a return, it changes the variable outside the scope
	fmt.Println(arrayCopy)
	fmt.Println(x)

}
