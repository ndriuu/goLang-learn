package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	andri := Man{"Andri"}
	andri.Married()
	fmt.Println(andri.Name)
}
