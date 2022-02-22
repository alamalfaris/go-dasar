package main

import "fmt"

//spamFilter bisa kita jadiin type lewat type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) { //Filter type dpt digunakan karena udh di declare as type
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("alam", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
