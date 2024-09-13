package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	FirstName string
	LastName  string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) SetLastName(lastName string) {
	p.LastName = lastName
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}

func main() {
	p := Person{}
	p.SetName("John Doe")
	p.SetAge(30)
	p.SetFirstName("John")
	p.SetLastName("Doe")
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
	fmt.Println(p.GetFirstName())
	fmt.Println(p.GetLastName())

	p2 := Person{Name: "Jane Doe", Age: 25, FirstName: "Jane", LastName: "Doe"}
	fmt.Println(p2.GetName())
	fmt.Println(p2.GetAge())
	fmt.Println(p2.GetFirstName())
	fmt.Println(p2.GetLastName())

}
