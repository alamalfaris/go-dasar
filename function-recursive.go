package main

import "fmt"

//recursive function itu bisa mangggil dirinya sendiri
//misal ada func search, didalemnya manggil search yg sama lagi

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	hasil := factorialRecursive(5)
	fmt.Println(hasil)
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
