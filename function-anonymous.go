package main

import "fmt"

func registerUser(name string, blacklist func(string) bool) {
	if blacklist(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//bisa bikin function nya kaya gini
	blacklist := func(name string) bool {
		return name == "alam"
	}

	registerUser("alam", blacklist)
	registerUser("wildan", blacklist)

	registerUser("odi", func(name string) bool {
		return name == "odi"
	})
}
