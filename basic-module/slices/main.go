package main

import "fmt"

func main() {

	// slices has 3 arguments in constructor to inicialize
	// ptr, len, cap
	// pointer, length and capacity
	mySlice := make([]int, 3, 6)
	// pointer, length = 3 and capacity = 6
	// this way, this variable declared at top will had this next value
	// [0|0|0| | | | ]
	//  0 1 2 3 4 5 6

	mySlice[1] = 9
	// [0|9|0| | | | ]
	//  0 1 2 3 4 5 6

	mySlice = append(mySlice, 2)
	// [0|9|0|2| | | ]
	//  0 1 2 3 4 5 6
	// pointer, length = 4 and capacity = 6

	anotherSlice := mySlice

	fmt.Println(mySlice)

	fmt.Println(anotherSlice)
	// here both will had the same value, but this 2 variables will appointed to the same pointer

	var slice []int

	// every slice what was declared without the function make inicialize nil (null)
	fmt.Println(slice)
	fmt.Println(slice == nil)

	var slice2 = []int{1, 2, 3}
	// slice2 := []int{1, 2, 3}
	// slice2 := make([]int, 3)

	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// when creating a slice, it's important, if possible, to declare it with the make function, 
	// because of the capacity attribute, which already "reserves" the space in memory, so it doesn't 
	// become a headache in the future

	testSlice := make([]int, 3, 4)

	testSlice[0] = 1
	testSlice[1] = 2
	testSlice[2] = 3
	testSlice = append(testSlice, 4)
	
	// at this moment testSlice is [1, 2, 3, 4]
	
	//testSlice[4] = 5
	// if i set index 5 in this slice, the application will 
	// return panic with index error, because the value setted is out of memory space reserved

	fmt.Println(testSlice)
	fmt.Println(cap(testSlice))
	
	testSlice = append(testSlice, 5)
	// but if i use append function, the array in a memory 
	// will be dropped and a new will be instancied in memory with a double of cap, in this case 8

	fmt.Println(testSlice)
	fmt.Println(cap(testSlice))

	testMemorySlice1 := make([]int, 2, 3)
	testMemorySlice2 := testMemorySlice1

	fmt.Println(testMemorySlice1)
	fmt.Println(testMemorySlice2)
	
	//testMemorySlice1[0] = 9
	
	// at this point the var testMemorySlice1 and testMemorySlice2 will be the same
	fmt.Println(testMemorySlice1)
	fmt.Println(testMemorySlice2)
	// [9, 0]
	// [9, 0]
	
	// but, if we append values it will "break" the cap definition
	testMemorySlice2 = append(
		testMemorySlice2, 
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	)

	// at this point, the variables don't have the same value because the variable
	// testMemorySlice broke the limit and initialized in a new memory space
	// thus forcing them to no longer have the same pointer, so when you change one, 
	// the other reference doesn't change anymore
	testMemorySlice2[0] = 9

	fmt.Println(testMemorySlice2)
	fmt.Println(testMemorySlice1) 
	// [9 0 1 2 3 4 5 6 7 8 9]
	// [9 0]

}