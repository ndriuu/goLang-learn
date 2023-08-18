package main

import "fmt"

func main() {
	fmt.Println("This is break")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Looping to ", i)
	}
	fmt.Println()
	fmt.Println("This is Continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Looping to ", i)
	}
}
