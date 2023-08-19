package main

import (
	"errors"
	"fmt"
)

func Division(number int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor cannot be zero")
	} else {
		result := number / divisor
		return result, nil
	}
}

func main() {
	result, err := Division(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
