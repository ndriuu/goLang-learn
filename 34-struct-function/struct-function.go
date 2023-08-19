package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var ndriuu Customer
	ndriuu.Name = "Ghufron Andriansyah"
	ndriuu.Address = "Indonesia"
	ndriuu.Age = 25
	ndriuu.sayHello("ALex")
}
