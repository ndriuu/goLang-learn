package main
import "fmt"
func main(){
	type numberId string //this way like create alias of data type, so wa can create variables concisely
	type Maried bool
	var numberIdme numberId = "0987654321"
	var mariedStatus Maried = true
	fmt.Println(numberIdme)
	fmt.Println(mariedStatus)
}