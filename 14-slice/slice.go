package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"Mey",
		"June",
		"Junly",
		"Agust",
		"September",
		"October",
		"NOvember",
		"December",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// months[5] = "Update"
	// fmt.Println(slice1)
	// slice1[0] = "Mey again"
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)
	var slice3 = append(slice2, "Andri")
	fmt.Println(slice3)
	slice3[1] = "Not December"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ghufron"
	newSlice[1] = "Andriansyah"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisArray := [5]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
