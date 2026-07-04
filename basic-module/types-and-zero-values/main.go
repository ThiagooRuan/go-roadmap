package main

import "fmt"

func main() {
	var (
		i     int
		f     float64
		s     string
		b     bool
		a     [3]int
		sl    []int
		m     map[string]int
		ch    chan int
		p     *int
		fn    func()
		iface interface{}
	)

	// primitive types
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte -> alias for uint 8
	// rune -> alias for int32
	// float32 float64
	// complex64 complex128

	fmt.Println("Zero Values")
	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("string: %q\n", s)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("array: %v\n", a)
	fmt.Printf("slice: %v\n", sl)
	fmt.Printf("map: %v\n", m)
	fmt.Printf("channel: %v\n", ch)
	fmt.Printf("pointer: %v\n", p)
	fmt.Printf("function: %v\n", fn)
	fmt.Printf("interface: %v\n", iface)

	//value with the format printed
	fmt.Printf("%T: %v\n", i, i)
}
