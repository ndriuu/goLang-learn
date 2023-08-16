package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ghufron",
		"address": "Tuban",
	}
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "goLang-Learn"
	book["author"] = "Andri"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
