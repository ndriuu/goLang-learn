package main

import "fmt"

func main() {
	result := getHello("Andri")
	fmt.Println(result)
	fmt.Println(getHello(""))
}

func getHello(name string) string {
	if name == "Andri" {
		return "Hello " + name
	} else {
		return "Who are you?"
	}
}
