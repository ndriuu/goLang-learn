package main

import "fmt"

func main() {
	//counter := 1
	// for counter <= 5 {
	// 	fmt.Println("Lopping to ", counter)
	// 	counter++
	// }

	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Looping to ", counter1)
	}

	slice := []string{"Ghufron", "Andri", "ansyah"}

	// for i := 0; i < len(slice); i++ {
	// 	if i == 1 {
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Print(slice[i])
	// }
	//fmt.Println()

	for i, value := range slice {
		fmt.Println("Index ", i, "= ", value)
	}
	for _, value := range slice { // code _ this, to tell the Go language that code is not necessary
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Ghufron"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
