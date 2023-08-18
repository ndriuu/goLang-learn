package main

import (
	"fmt"
)

func main() {
	name := "Ghufron"
	switch name {
	case "Ghufron":
		fmt.Println("Helo Ghufron")
		fmt.Println("Morning Ghufron")
	case "Andri":
		fmt.Println("Hello Andri")
		fmt.Println("How are you, Andri?")
	default:
		fmt.Println("hi, who are you?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("the name is too long")
	// case false:
	// 	fmt.Println("the name is right")
	// }

	// Another way
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("The name is too long")
	case length > 5:
		fmt.Println("The name is a bit long")
	default:
		fmt.Println("the name is right")
	}
}
