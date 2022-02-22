package main

import "fmt"

func main() {
	result := 10 + 10
	fmt.Println(result)

	result += 10
	fmt.Println(result)

	result++
	fmt.Println(result)

	isChecked := false

	if !isChecked {
		fmt.Println("isChecked = ", isChecked)
	}

	x := -1
	fmt.Println(x)
}
