package main

import "fmt"

func main() {
	var exam = 81
	var attendance = 81

	var passExam = exam > 80
	var passAttendace = attendance > 80
	fmt.Println(passExam)
	fmt.Println(passAttendace)

	var pass = passExam && passAttendace
	fmt.Println(pass)

	fmt.Println(exam > 80 && attendance > 80) //usually like this
}
