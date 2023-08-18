package main

import "fmt"

func getFullName() (firstName string, lastName string, instagramName string) {
	firstName = "Ghufron"
	lastName = "Andriansyah"
	instagramName = "ndriu._"
	return
}
func main() {
	firstName, lastName, instagramName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(instagramName)
}
