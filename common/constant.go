package common

import "fmt"

const (
	Alam = "Alam Alfaris"
	ardi = "Ardiatama F" // karna cons ini diawali huruf kecil, jadi gak bisa di akses di package lain
)

func Pergi() {
	fmt.Println("lagi pergi nih")
}

func pulang() { // karna function ini diawali huruf kecil, jadi gak bisa di akses di package lain
	fmt.Println("baru mau pulang")
}
