package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Print() {
	fmt.Printf("%+v\n", u)
}

// To understand the concept of pointers, we first need to talk about memory
// each allocated variable is like a drawer in memory, and each drawer has an address
// something like 0x100, 0x101, 0x102

// when we define age := 24 for example, something like this happens
// address | value
// 0x100   | 24

// the pointer doesn't store the value, it actually points to where the value is, which drawer it's in

func main() {
	// declaring a variable
	age := 24

	// The & symbol points to the memory address.
	// by using it, I'm saying I want to store the memory address of this
	// age variable in the p variable
	p := &age

	fmt.Println(age, p)
	//24 0x37f71240a0

	// The * symbol has two meanings
	// The first is that when declaring a variable - var p *int, it declares that the variable is a pointer
	// in this example, it means p is a pointer to int
	// The second meaning is that it also refers to the value stored in the slot, for example age2 := *p
	// the variable p holds the memory address, which points to the slot that has the value, using age2 := *p,
	// I get the value that is in the slot using the address provided by the pointer
	// The * operator dereferences the pointer,
	// allowing access to the value stored at that address.

	fmt.Println(age)  // 24
	fmt.Println(&age) // 0x37f71240a0
	fmt.Println(*p)   // 24

	// changing a variable through the pointer
	*p = 30
	// *p = 30 basically means, "I want the value stored at that address to become 30"

	fmt.Println(age)

	// Nil Pointer
	// A pointer might not point to anything
	var p2 *int

	fmt.Println(p2) //<nil>
	// fmt.Println(*p2) //panic, because it is pointing to a non-existent memory space

	p2 = &age

	fmt.Println(p2)
	// A pointer is just another variable.
	// The value stored inside it is a memory address.

	// Pointers don't copy data from the objects they point to; they point to the variable that's already allocated with the value.
	// Imagine a scenario where you have a struct with 500 fields, for example. In this case, when I wanted to save the data, if
	// I didn't pass them as pointers, every time I ran the save function, I would allocate a copy of this struct in memory. By using pointers,
	// I reference the same memory space, which is why pointers are used—for performance, and also because they allow direct modifications.

	user := User{
		Name: "Thiago",
		Age:  24,
	}

	user.Print()

	// two pointers were created pointing to the same memory space, if I change the value using one of the references, I adjust the value that both pointers are pointing to,
	// this proves that they refer to the same memory space
	p3 := &user.Age
	p4 := p3

	*p4 = 100
	user.Print()
}
