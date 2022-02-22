package main

import "fmt"

func main() {
	type nama string
	type nikah bool

	var namaAlam nama = "Alam Alfaris"
	var statusNikah nikah = false

	fmt.Println("nama = ", namaAlam)
	fmt.Println("sudah nikah = ", statusNikah)
}
