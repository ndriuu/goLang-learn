package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Ghufron")
	data.PushBack("Andriansyah")
	data.PushBack("ndriuu")
	data.PushFront("Name")

	// The order is from the front
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	//The order is from the back
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
