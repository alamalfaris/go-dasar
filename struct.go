package main

import "fmt"

//struct ini kalo di oop seperti class model

type Customers struct {
	Name    string
	Address string
	Age     int
}

// struct method
// method biasa (sebetulnya) tapi seolah2 ada didalam struct itu
// bentuk function nya gini
func (cus Customers) sayHi() {
	fmt.Println("Hi", cus.Name, "aku tinggal di", cus.Address)
}

func main() {
	// cara declare struct-1
	var alam Customers
	alam.Name = "alam"
	alam.Address = "jakarta"
	alam.Age = 24

	alam.sayHi()

	// fmt.Println(alam)
	// fmt.Println(alam.Name)
	// fmt.Println(alam.Address)
	// fmt.Println(alam.Age)

	// // cara declare struct-2
	// budi := Customers{
	// 	Name:    "budi",
	// 	Address: "bogor",
	// 	Age:     21,
	// }
	// fmt.Println(budi)

	// // cara declare struct-3
	// indah := Customers{"indah", "tangerang selatan", 22} // ini urutannya harus sama, sesuai model nya
	// fmt.Println(indah)

}
