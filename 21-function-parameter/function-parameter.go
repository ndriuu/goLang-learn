package main

import "fmt"

func main() {
	sayHelloTo("Ghufron", "Andriansyah")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
