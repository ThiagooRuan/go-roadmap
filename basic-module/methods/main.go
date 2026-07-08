package main

import "fmt"

// methods are functions attached to the type

type User struct {
	Name string
	Age  int
}

type Car struct {
	Brand string
	Model string
	Year  int
}

// instead of
// func printUser(user User){
// ...
// }

func (u User) Print() {
	fmt.Printf("%+v\n", u)
}

// the part (u User) is called the receiver, responsible for
// declaring that this method belongs to the User type. The receiver is a variable,
// but the good convention is to use a single letter, for example
// func (u User)
// func (c Car)
// func (a Account)

// The receiver defines which type owns the method.
// Value receiver
// in this way, this method only gets a copy of the struct,
// so even if you change it, it will only be inside the function's scope
// to change not the copy but the reference, the correct way would be to use a pointer receiver
// func (u User) ChangeName() {
// u.Name = "Carlos"
// fmt.Printf("%+vn", u)
// }

// Pointer Receiver
// this way it points to the reference and doesn’t copy the struct into a new memory space
// when initializing the struct, instead it modifies the same memory space of the reference,
// pointer receiver should be used when the struct is being modified or when it’s large,
// it avoids copying the whole struct on each method call.
func (u *User) ChangeName(name string) {
	u.Name = name
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (c *Car) UpdateBrand(brand string) {
	c.Brand = brand
}

func (c *Car) UpdateModel(model string) {
	c.Model = model
}

func (c *Car) UpdateYear(year int) {
	c.Year = year
}

func (c *Car) Print() {
	fmt.Printf("%+v\n", c)
}

func main() {
	user := User{
		Name: "Thiago",
		Age:  24,
	}

	user.Print()
	// when we use a method with a pointer receiver, it's not necessary to use &
	// to reference the same memory space, Go already does this natively, taking the address when needed

	user.ChangeName("Carlos")
	// (&user).ChangeName()

	fmt.Println("User is Adult? ", user.IsAdult())
	user.Print()

	// ==================
	// CAR Examples
	// ==================

	car := Car{
		Brand: "VW",
		Model: "Golf",
		Year:  2020,
	}

	car.Print()

	car.UpdateBrand("Chevy")
	car.Print()

	car.UpdateModel("Jetta")
	car.Print()

	car.UpdateYear(2010)
	car.Print()

	// Methods are not classes, in Java for example a class has attributes and has methods tied to it,
	// in Go the struct exists, the method is created and tied to it, but the struct is still just data

	// Dealing with methods, not everything needs to be by reference (pointer),
	// but also not everything needs to be by value (copy)

	// Mixing receivers is allowed, but usually projects that have methods that change state choose
	// to use pointer receivers consistently for all the struct's methods
}
