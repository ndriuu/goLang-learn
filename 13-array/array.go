package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ghufron"
	names[1] = "Andri"
	names[2] = "ansyah"
	//names[3]	= "ansyah"
	//This will be an error because the array capacity of the array variable is only up to index 2 or only enough to hold 3 elements.

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//another way
	var values = [5]int{
		90,
		95,
		85,
	}
	fmt.Println(values)
	fmt.Println(values[1])
	fmt.Println(len(names))
	fmt.Println(len(values))
}
