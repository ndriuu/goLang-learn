package main

import "fmt"

func main() {
	name := "Andri"
	if name == "Ghufron" {
		fmt.Println("Hello World")
	} else if name == "Andri" {
		fmt.Println("Hi, what's your name?")
	} else {
		fmt.Println("How are you?")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("Too Long")
	} else {
		fmt.Println("The name is right")
	}
}
