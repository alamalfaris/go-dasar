package main

import "fmt"

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	firstName, lastName := getFullName("alam", "alfaris")
	fmt.Println(firstName)
	fmt.Println(lastName)

	namaAwal, _ := getFullName("alam", "alfaris") //underscore untuk menghiraukan lastName
	fmt.Println(namaAwal)
}
