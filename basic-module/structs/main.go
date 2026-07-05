package main

import "fmt"

// A struct is a composite data type that groups related values.
type User struct {
	Name string
	Age  int
}

type Address struct {
	Country string
	City    string
}

type Employee struct {
	Name    string
	Address Address
}

type Car struct {
	Brand string
	Model string
	Year  int
}

func printCar(car Car) {
	fmt.Printf("%+v\n", car)
}

func updateYear(car *Car) {
	car.Year = 2000
}

func changeBrand(car *Car) {
	car.Brand = "Chevy"
}

func printUser(user User) {
	fmt.Printf("%+v\n", user)
}

func modifyUser(user *User) {
	user.Name = "Carlos"
}

func main() {
	// ======================
	// User Examples
	// ======================

	// also being able to be defined this way
	var userData User

	fmt.Println(userData) // zero values of user struct { 0}

	userData.Name = "Thiago"
	userData.Age = 24

	user := User{
		Name: "Thiago",
		Age:  24,
	}

	printUser(user)
	// if i want modify user name, I would have to use a pointer
	// User is passed by value.
	// modifyUser receives a pointer so it can modify
	// the original struct instead of a copy.
	modifyUser(&user)

	printUser(user)

	// ======================
	// Employee Examples
	// ======================

	employee := Employee{
		Name: "John",
		Address: Address{
			Country: "USA",
			City:    "New york",
		},
	}

	fmt.Println(employee)

	// ======================
	// Car Examples
	// ======================

	car := Car{
		Brand: "VW",
		Model: "Fox",
		Year:  2010,
	}

	car2 := Car{
		Brand: "VW",
		Model: "Fox",
		Year:  2010,
	}

	fmt.Println("car its equal to car2? ", car == car2)

	printCar(car)

	updateYear(&car)
	printCar(car)

	changeBrand(&car)
	printCar(car)

	fmt.Println("car its equal to car2? ", car == car2)

	car2.Brand = "Chevy"
	car2.Year = 2000

	fmt.Println("car its equal to car2? ", car == car2)
}
