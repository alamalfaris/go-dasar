package main

import "fmt"

func powerAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	total := powerAll(2, 3, 4)
	fmt.Println(total)
}
