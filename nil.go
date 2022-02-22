package main

import "fmt"

//di bahasa lain kalo bikin variable tanpa inisialisasi data, maka dia defaultnya null
//kalo di golang, dia akan bernilai default tipe datanya misal int = 0; string = "", bool = false
//nil = null
//nil bisa digunakan untuk validasi
//nil cuma bisa digunakan di beberapa tipe data :
//interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil //nil disini digunakan untuk function
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("alam")
	fmt.Println(person)

	var andi map[string]string
	if andi == nil { // nil disini digunakan untuk map
		fmt.Println("andi kosong")
	} else {
		fmt.Println("andi ada isinya")
	}
}
