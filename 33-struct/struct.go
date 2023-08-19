package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ndriuu Customer
	ndriuu.Name = "Ghufron Andriansyah"
	ndriuu.Address = "Indonesia"
	ndriuu.Age = 25

	fmt.Println(ndriuu.Name)
	fmt.Println(ndriuu.Address)
	fmt.Println(ndriuu.Age)

	// Another way and suggest use this
	alex := Customer{
		Name:    "Alexander",
		Address: "Australia",
		Age:     30,
	}
	fmt.Println(alex)

	// Another way and not suggest use this cuz it is error prone
	speed := Customer{"IshowSpeed", "England", 40}
	fmt.Println(speed)

}
