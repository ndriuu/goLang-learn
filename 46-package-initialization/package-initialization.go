package main

import (
	"fmt"
	"goLang-Learn/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
