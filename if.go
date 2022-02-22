package main

import "fmt"

func main() {
	// umur := 10
	// gender := "wanita"

	// if umur >= 20 && gender == "pria" {
	// 	fmt.Println("pria diatas 20")
	// } else if umur < 20 && gender == "pria" {
	// 	fmt.Println("pria dibawah 20")
	// } else {
	// 	fmt.Println("zzz")
	// }

	nama := "alam alfaris kusumanto putra"

	if len(nama) > 10 {
		fmt.Println("nama kepanjangan")
	} else {
		fmt.Println("nama udah aman")
	}

}
