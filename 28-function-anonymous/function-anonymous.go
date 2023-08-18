package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're bloked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }
// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Andri", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Andri", func(name string) bool {
		return name == "root"
	})
}
