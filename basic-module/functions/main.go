package main

import (
	"errors"
	"fmt"
)

// func -> function declare
// sum( -> Identifier
// a, b int -> params type definition - since it's the same type i can define it just once with int
// )
// int -> return type
// { return a + b } -> function body
func sum(a, b int) int {
	return a + b
}

// functions with no return
// no need to specify the return type
func logDebug(msg string) {
	fmt.Printf("[debug] '%s'\n", msg)
}

// multiple returns
func div(a, b int) (int, int) {
	var (
		quotient  = a / b
		remainder = a % b
	)

	return quotient, remainder
}

// multiple returns principal utility
func floatDiv(c, d float64) (float64, error) {
	if d == 0 {
		return 0, errors.New("Division by zero")
	}

	return c / d, nil
}

// named returns
func calculateRectangleParams(
	width, height float64,
) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)

	return area, perimeter
}

// naked return
func nakedCalculateRectangleParams(
	width, height float64,
) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)

	return
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// variadic parameter -> You can only have one variadic parameter and it has to be at the end of the function declaration
func sumAll(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func printWithPrefix(prefix string, msgs ...string) {
	for _, msg := range msgs {
		fmt.Println(prefix, msg)
	}
}

func main() {
	a := 5
	b := 3

	logDebug("Begin")
	fmt.Println(sum(a, b))

	// multiple assignments
	quotient, remainder := div(a, b)

	fmt.Printf("dividing %v by %v\nquotient = %v\nremainder = %v\n", a, b, quotient, remainder)

	println(floatDiv(5.2, 2.0))

	// anonymous functions, functions that don’t have an identifier
	mult := func(a, b int) int {
		return a * b
	}

	fmt.Println(mult(4, 5))

	//assigning a function to a variable
	sum := add

	fmt.Println(sum(1, 2))

	// Passing a function as an argument for a function to execute
	fmt.Println(operate(2, 4, sum))
	fmt.Println(operate(10, 2, sub))

	// setting a function inside a variable, passing a value, and then executing the function with the value with 5 turning it into 10
	double := getMultiplier(2)
	fmt.Println(double(5))

	nums := []int{1, 3, 2, 6}
	fmt.Println(sumAll(nums...))

	msgs := []string{"bar", "baz", "qux"}
	printWithPrefix("foo", msgs...)

	logDebug("End of the execution")
}
