package main
import "fmt"
func main(){
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 83
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	
	i++ // i = i + 1
	fmt.Println(i)
	
	var negative = -100
	var positive = +100 //this is not mandatory or necessary, because even without positive 100 is already a positive number
	fmt.Println(negative)
	fmt.Println(positive)
}