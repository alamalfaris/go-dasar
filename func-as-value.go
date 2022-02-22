package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("alam")
	fmt.Println(result)
	fmt.Println(sayGoodBye("alam"))
	fmt.Println(getGoodBye("alam"))
}
