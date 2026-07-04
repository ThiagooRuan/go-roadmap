package main

import "fmt"

func main() {
	fmt.Println("Iniciando o programa...")

	// defer was executed anyway
	defer fmt.Println("Defer é executado mesmo assim!!")

	// when the panic function is called the expected behavior is an exception
	panic("Algo deu errado!!")

	// when in doubt, don't use panic -> isn't a good convention

	fmt.Println("Isso nunca será executado")
}