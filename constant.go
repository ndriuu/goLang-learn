package main
import "fmt"
func main(){
	const firsname = "Ghufron"
	const lastname = "Andriansyah"
	const age = 25

	//fmt.Println(firsname)
	fmt.Println(lastname)
	fmt.Println(age)	
	//note : constants are different from variables, variables will become errors if not used, but constants are just the opposite

	// Another way : create more than 1 constant
	const(
		country string 	= "Indonesia"
		city			= "Tuban"
		score 			= 99.9
	)
	fmt.Println(country)
	fmt.Println(city)
	fmt.Println(score)	
}