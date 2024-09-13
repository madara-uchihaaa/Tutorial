package main

import (
	"fmt"
	"sort"
)

func main() {
	// Array in Golang

	var x [5]int
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	// x[5] = 60 // error: invalid array index 5 (out of bounds for 5-element array)

	fmt.Println(len(x))

	var fruits [5]string = [5]string{"apple", "banana", "mango", "orange", "grapes"}

	fmt.Println(fruits)

	// Slices in Golang

	fmt.Println("Welcome to Slices in Golang")

	var students = []string{"John", "Doe", "Smith", "Alex", "Ram"}

	students = append(students, "Raj", "Rahul")

	fmt.Println(students)

	students = append(students[1:3], students[4:]...)
	// O/P: [Doe Smith Ram Raj Rahul]

	fmt.Println(students)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 458
	highScores[3] = 234

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"ReactJS", "NodeJS", "MongoDB", "Python", "Django", "Flask"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	// ... is used to append all the elements of the slice

	fmt.Println(courses)

}
