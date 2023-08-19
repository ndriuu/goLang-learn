package main

import "fmt"

//============ Defer ============
// func main() {
// 	runApplication(0)
// }
// func logging() {
// 	fmt.Println("Finished calling the function")
// }
// func runApplication(value int) {
// 	defer logging()
// 	fmt.Println("Run application")
// 	result := 10 / value
// 	fmt.Println("Result", result)
// }

//============ Panic ============
// func main() {
// 	runApp(false)
// }
// func endApp() {
// 	fmt.Println("Application is complete")
// }
// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("Application is error")
// 	}
// 	fmt.Println("Application running")
// }

// ============ Recover ============
func main() {
	runApp(true)
	fmt.Println("Andri")
}
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error with message:", message)
	}
	fmt.Println("Application is complete")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application is error")
	}
	fmt.Println("Application running")
}
