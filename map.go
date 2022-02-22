package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "alam",
		"address": "senayan",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["role"] = "programmer" //nambah key role dgn value programmer
	fmt.Println(person)
	fmt.Println(len(person))

	delete(person, "role") // hapus dgn key

	fmt.Println(person)

	person["name"] = "eko" // ubah value di key name jadi eko
	fmt.Println(person)

	//bikin map dari awal
	student := make(map[string]string)
	student["name"] = "andi"
	student["grade"] = "A"
	fmt.Println(student)
}
