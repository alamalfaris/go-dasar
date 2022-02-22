package main

import "fmt"

func main() {
	var months = [...]string{ //... tu si len array gak didefine didepan
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}
	fmt.Println(months)
	fmt.Println(months[0:3])

	var slice1 = months[1:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "rrr"
	fmt.Println(slice1)
	fmt.Println(months)

	slice1 = append(slice1, "zzz")
	//fmt.Println(slice2)
	fmt.Println(slice1)
	fmt.Println(months)

	var slice2 = make([]string, 2, 5) //buat slice dgn make | 2 tu length | 5 tu capacity
	slice2[0] = "alam"
	slice2[1] = "alfaris"
	fmt.Println(slice2)

	var slice3 = make([]string, len(slice2), cap(slice2))
	copy(slice3, slice2)
	fmt.Println(slice3)

	// hati2 saat bikin array dan slice, beda tipis
	// bikin array
	iniArray := [...]int{1, 2, 3, 4, 5}
	//bikin slice
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
