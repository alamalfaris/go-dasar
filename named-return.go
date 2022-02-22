package main

import "fmt"

func namedReturn() (first string, second string) { //first dan second adalah variable untuk return
	first = "alam"
	second = "alfaris"
	return
}

func main() {
	a, b := namedReturn()
	fmt.Println(a)
	fmt.Println(b)
}
