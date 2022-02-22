package main

import "fmt"

//recover adalah function yang bisa kita gunakan untuk menangkap data panic
//dengan recover proses panic akan terhenti, sehingga program tetap berjalan

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error:", message)
	}
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
	fmt.Println("alam")
}
