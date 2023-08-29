package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("y([a-z])u")

	fmt.Println(regex.MatchString("you"))
	fmt.Println(regex.MatchString("yzu"))
	fmt.Println(regex.MatchString("yOu"))

	search := regex.FindAllString("you yuu yii yaa yee", -1)
	fmt.Println(search)

}
