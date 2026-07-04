package main

import "fmt"

// Using defer is gonna be delay the execution of this expression, 
// that is, the result here will be world!!! hello
func helloWorld() {
	defer fmt.Println("Hello")
	fmt.Println("World!!!")
}

// defer respect the order lifo (last in first out) to execute, this way, the result here will be 1,2,3,4
func printNumbers() {
	defer fmt.Println("4")
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1")
}

// even if defer print will run after the normal print line, 
// at the moment that defer print line will run it's result 10 
// because, the moment defer was called the variable was passed with the value 10
func varAssingmentTest() {
	x := 10
	defer fmt.Println("Valor de x no defer: ", x)
	
	x = 20
	fmt.Println("Valor de x no defer: ", x)
}


func getValue() int {
	fmt.Println("chamando getValue()")
	return 42
}

func main() {
	helloWorld()
	printNumbers()
	varAssingmentTest()

	// the order of this execution results in
	// chamando getValue()
	// Função main retornando
	// Valor obtido:  42
	// even though the value has been setted, the print only will print only when the main function return
	defer fmt.Println("Valor obtido: ", getValue())
	fmt.Println("Função main retornando")
}

