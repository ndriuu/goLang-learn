package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init is Called")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
