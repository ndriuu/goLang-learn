package main

import "fmt"

type Hasname interface {
	Getname() string
}

func SayHello(hasName Hasname) {
	fmt.Println("Hello", hasName.Getname())
}

func (person Person) Getname() string {
	return person.Name
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) Getname() string {
	return animal.Name
}

func main() {
	var andri Person
	andri.Name = "Ghufron Andriansyah"
	SayHello(andri)

	cat := Animal{
		Name: "Meongg",
	}
	SayHello(cat)
}
