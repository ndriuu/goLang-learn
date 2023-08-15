package main
import "fmt"
func main(){
	var number32 int32 = 100000
	var number64 int64 = int64(number32)
	var number8 int8 = int8(number32) 

	fmt.Println(number32)
	fmt.Println(number64)
	fmt.Println(number8)// this result would be strange because the maximum value of int8 does not reach the value of the number32 variable. 

	var name = "Ghufron"
	var g byte = name[0] // byte is alias of uint8
	var gString string = string(g) // this is a conversion of string

	fmt.Println(name)
	fmt.Println(gString)
	
}