package main

import "fmt"

func main() {
	// declaring a slice
	a := []int{1, 2, 3, 4, 5}
	// ptr, len = 5, cap = 5
	// [1|2|3|4|5]
	//  0 1 2 3 4

	// this slice recieve a slice expression
	s := a[1:4]
	// ptr, len = 3, cap = 4
	// cap = 4 because the underlying array goes to the end of the original array, even without receiving these values
	// [2|3|4| ]
	//  1 2 3

	fmt.Println(s)

	// I can also create a slice expression like this,
	// leaving out the high or low values, or even not setting
	// either high or low, which returns the same array
	x := a[2:]
	fmt.Println(x)

	y := a[:2]
	fmt.Println(y)

	z := a[:]
	fmt.Println(z)

	// I can also ´"delete" a element from a slice using
	// slice_expression, creating a new slice without a element
	b := []int{1, 2, 3, 4, 5}

	c := append(b[:2], b[3:]...) // using append this way destructur the subslice and add one a one
	// c := append(b[:2], b[3], b[4]) -- return the same thing
	fmt.Println(c)
	// result [1 2 4 5]

	//creating a slice_expression from array
	d := [5]int{1, 2, 3, 4, 5}
	e := d[:2] // but, at this point, this variable is a slice, not a array
	fmt.Println(e)

	// I can create a slice_expression from string
	str := "hello"
	subStr := str[1:4] // result ell
	
	fmt.Println(subStr)
}
