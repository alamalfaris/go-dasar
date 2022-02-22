package main

import "fmt"

func main() {
	// var age int = 1
	// fmt.Println(age)
	fmt.Println("enter your first name: ")
	var firstName string
	fmt.Scanln(&firstName)

	fmt.Println("enter your last name: ")
	var lastName string
	fmt.Scanln(&lastName)

	fmt.Println("your full name is: " + firstName + " " + lastName)
}
