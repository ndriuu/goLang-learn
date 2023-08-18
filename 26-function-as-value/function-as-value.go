package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("Andri")
	fmt.Println(result)
	fmt.Println(getGoodBye("Andri"))
}
func getGoodBye(name string) string {
	return "Good Bye " + name
}
