package main

import "fmt"

//panic
//- panic function adalah function yang bisa digunakan untuk menghentikan program
//- panic function biasanya dipanggil saat program kita ada error
//- saat panic func dipanggil, program akan terhenti, tapi defer func tetap dieksekusi

func endApp() {
	fmt.Println("APLIKASI SELESAI")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("APLIKASI BERJALAN")
}

func main() {
	runApp(true)
}
