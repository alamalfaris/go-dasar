package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "alam"
	names[1] = "alfaris"
	names[2] = "putra"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[0] = "endi"
	fmt.Println(names[0])

	var angka = [3]int{ //declare array with direct values
		30, 20, 50,
	}
	fmt.Println(angka)

	fmt.Println(len(names))
	fmt.Println(len(angka))
}
