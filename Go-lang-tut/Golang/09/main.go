package main

func main() {
	// If else in Golang

	age := 18

	if age >= 18 {
		println("You can vote")
	} else if age >= 16 {
		println("You can drive")
	} else {
		println("You can have fun")
	}

	if num := 3; num == 3 {
		println("Num is 3")
	} else if num == 4 {
		println("Num is 4")
	} else {
		println("Num is not 3 or 4")
	}

	// Switch in Golang

	switch value := 20; value {
	case 16:
		println("You can drive")
	case 18:
		println("You can vote")
	default:
		println("You can have fun")
	}

	// fallthrough in Golang means that if the case is true, it will execute the next case as well
}
