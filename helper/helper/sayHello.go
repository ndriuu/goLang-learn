package helper

import "fmt"

var version = 1
var Application = "goLang Learn"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// can not be accessed outside cuz the first letter is not capitalized
func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
