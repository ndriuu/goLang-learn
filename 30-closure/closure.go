package main

import "fmt"

func main() {
	name := "Ghufron"
	counter := 0

	increment := func() {
		name = "Andriansyah"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
