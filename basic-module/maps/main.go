package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// A map is a data structure that stores key-value pairs.
// It is useful when you need to retrieve a value using a key.
// Examples:
// CPF   -> Customer
// ID    -> Product
// Email -> User
// Looking up a value in a map is much faster than iterating through a slice.

func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func main() {

	// Another way to initialize a map.
	// ages := make(map[string]int)

	// ages["Thiago"] = 24
	// ages["Carlos"] = 30
	// ages["John"] = 35

	ages := map[string]int{
		"Thiago": 24,
		"Carlos": 30,
		"John":   35,
	}

	printMap(ages)

	fmt.Println(ages["Thiago"]) // Accessing a value using its key.

	ages["Thiago"] = 40         // Updating the value associated with the key.
	fmt.Println(ages["Thiago"]) // 40

	// Assigning a value to a new key automatically inserts it into the map.
	ages["Jorge"] = 42

	printMap(ages)

	delete(ages, "Carlos") // Removes the entry associated with the key.
	printMap(ages)

	ages["Pedro"] = 8

	// The second return value indicates whether the key exists.
	age, exists := ages["Pedro"]

	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("User not found")
	}

	// Maps are reference types.
	// Unlike structs, assigning one map to another does not create a copy.
	// Instead, both variables share the same underlying data.
	ages2 := ages

	fmt.Println(ages, ages2)

	// Updating one map also affects the other.
	ages2["Thiago"] = 18

	fmt.Println(ages, ages2)

	// Structs can be compared using ==.
	// Maps cannot be compared with each other.
	// The only valid comparison is against nil.
	// ages == ages2 // Compilation error.

	if ages != nil {
		fmt.Println("Ages is not nil")
	}

	users := make(map[int]User)

	users[1] = User{
		Name: "Thiago",
		Age:  20,
	}

	users[2] = User{
		Name: "Carlos",
		Age:  22,
	}

	users[3] = User{
		Name: "John",
		Age:  40,
	}

	fmt.Println(users)
	fmt.Println(users[1])

	users[1] = User{
		Name: "Diogo",
		Age:  30,
	}

	fmt.Println(users)

	delete(users, 3)

	fmt.Println(users)

	for id, user := range users {
		fmt.Println(id, user)
	}

	users2 := users

	fmt.Println(users)
	fmt.Println(users2)

	// Since both variables share the same underlying data,
	// updating one map also updates the other.
	users2[1] = User{
		Name: "Thiago",
		Age:  24,
	}

	fmt.Println(users)
	fmt.Println(users2)

}
