package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ghufron Andriansyah", "Andri"))
	fmt.Println(strings.Contains("Ghufron Andriansyah", "Alexader"))

	fmt.Println(strings.Split("Ghufron Andriansyah", " "))

	fmt.Println(strings.ToLower("JAPAN"))
	fmt.Println(strings.ToUpper("indonesia"))
	fmt.Println(strings.ToTitle("ghufron andriansyah"))
	fmt.Println(strings.Trim("      Kitten   a     ", " "))
	fmt.Println(strings.ReplaceAll("Ghufron ndriu._ Ghufron Ghufron", "Ghufron", "Andri"))

}
