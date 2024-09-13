package main

import "fmt"

// jwt := "some jwt"
// error: non-declaration statement outside function body
// This is because we can only use the := syntax inside a function body.

// var jwt = "some jwt" // works


const LoginToken string = "asdasdasdasdasd"
// L is capital because we want to export it
// const loginToken string = "asdasdasdasdasd"
// L is small because we don't want to export it

func main() {

	// Variables

	fmt.Println("Hello World")

	var x string = "Hello World"

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var isLogged bool = true

	fmt.Println(isLogged)
	fmt.Printf("%T\n", isLogged)

	var y uint8 = 255
	// y = 256 // error: constant 256 overflows uint8

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	var z float32 = 3.14777727272727727272727272727

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	var a float64 = 3.14777727272727727272727272727

	fmt.Println(a)
	fmt.Printf("%T\n", a)


	var b complex64 = 1 + 2i

	fmt.Println(b)
	fmt.Printf("%T\n", b)


	// alias and default types

	var val int
	// Default value is 0

	fmt.Println(val)
	fmt.Printf("%T\n", val)

	var val2 string
	// Default value is ""

	fmt.Println(val2)
	fmt.Printf("%T\n", val2)

	// implicit type

	var val3 = "Hello World"

	fmt.Println(val3)
	fmt.Printf("%T\n", val3)

	// no var style

	val4 := "Hello World"

	fmt.Println(val4)
	fmt.Printf("%T\n", val4)

	// multiple declaration

	var (
		name string = "John"
		age int = 35
		height int32 = 72
		isMarried bool = false
	)

	fmt.Println(name)
	fmt.Printf("%T\n", name)
	fmt.Println(age)
	fmt.Printf("%T\n", age)
	fmt.Println(height)
	fmt.Printf("%T\n", height)
	fmt.Println(isMarried)
	fmt.Printf("%T\n", isMarried)


	fmt.Println(LoginToken)
}
