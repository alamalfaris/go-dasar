package main

import "fmt"

func tambah(num1 int, num2 int) int {
	return num1 + num2
}

func kurang(num1 int, num2 int) int {
	return num1 - num2
}

func kali(num1 int, num2 int) int {
	return num1 * num2
}

func bagi(num1 int, num2 int) int {
	return num1 / num2
}

func myCalculator(num1 int, num2 int, filter func(int, int) int) {
	hasil := filter(num1, num2)
	fmt.Println(hasil)
}

func main() {
	myCalculator(1, 2, tambah)
	myCalculator(10, 5, kurang)
	myCalculator(10, 10, kali)
	myCalculator(6, 3, bagi)
}
