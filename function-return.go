package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Halo Bro"
	} else {
		return "Halo " + name
	}
}

func main() {
	fmt.Println(getHello("alam"))
	fmt.Println(getHello(""))
}
