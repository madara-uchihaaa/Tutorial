package main

import "fmt"

func main() {
	// new() function vs make() function

	// new() function
	// new() function is used to allocate memory to a variable.
	// It takes a type as an argument and returns a pointer to a newly allocated zero value of the type passed as an argument.
	// The zero value is the value that is the default value of the type.
	// For example, the zero value of int is 0, the zero value of string is an empty string, the zero value of bool is false, and so on.
	// new() function is not used to initialize a variable.
	// It is used to allocate memory to a variable only.


	// make() function
	// make() function is used to create a slice, map, or channel.
	// It takes a type and a length as an argument and returns an initialized slice, map, or channel.
	// The length argument is optional and is only required when creating a slice, map, or channel without using a composite literal.
	// For example, make([]int, 10) returns a slice of length 10 with all elements initialized to their zero value.
	// make() function is not used to allocate memory to a variable.
	// It is used to create a slice, map, or channel only.


	fmt.Println("Welcome to a class of Pointers")

	var x int = 10
	var intPtr *int = &x

	// &x is the address of x, reference means &x
	// *int is the type of intPtr means it is a pointer to an int type variable

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)

	fmt.Println("Value of intPtr:", intPtr)
	fmt.Println("Address of intPtr:", &intPtr)
	fmt.Println("Value at address intPtr:", *intPtr)

	// *intPtr is the value at the address stored in intPtr, dereference means *intPtr

	*intPtr = 20

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)

	fmt.Println("Value of intPtr:", intPtr)
	fmt.Println("Address of intPtr:", &intPtr)
	fmt.Println("Value at address intPtr:", *intPtr)

}