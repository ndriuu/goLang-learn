package main

import (
	"fmt"
	"goLang-Learn/helper/helper"
)

func main() {
	helper.SayHello("Andri")
	// helper.sayGoodbye("Andri") // Error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // Error
}
