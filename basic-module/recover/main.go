package main

import "fmt"

func main() {
	//recover is used to gracefully terminate the program, often to free resources,
	// and when nested with panic, it prevents further use of these resources, since
	// even if an error occurs, the defer code will still be executed.
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("Recuperado do panic: ", r)
		}
	}()

	fmt.Println("Iniciando execução...")
	panic("Algo deu errado!!")
	fmt.Println("Isso nunca será executado")

}
