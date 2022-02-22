package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Alam Alfaris", "Alam")) //cek apakah "Alam" ada di dalam "Alam Alfaris"
	fmt.Println(strings.Contains("Alam Alfaris", "Ardi"))

	fmt.Println(strings.Split("Alam Alfaris", " ")) //split string berdasarkan separator

	fmt.Println(strings.ToLower("Alam Alfaris"))
	fmt.Println(strings.ToUpper("Alam Alfaris"))

	fmt.Println(strings.Trim("a   Alam Alfaris   a", " "))    //utk hapus spasi di kiri dan kanan, tapi kalo ada huruf diujung nya jadi gak bisa
	fmt.Println(strings.ReplaceAll("Alam Alfaris", "A", "4")) // utk replace semua "A" di "Alam Alfaris" dgn "4"
}
