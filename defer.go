package main

import "fmt"

//defer itu fitur untuk memanggil function y setelah function x dijalankan
// tanpa peduli function x nya error atau sukses

func logging() {
	fmt.Println("ini func logging")
}

func runApp(value int) {
	defer logging() //nah pake defer ini, biar tetep jalanin func logging setelah runApp
	//taruh nya di paling atas karena biar gak kena error di bawah (kalo ada error)
	fmt.Println("ini func runApp")
	hasil := 10 / value
	//logging() //biasanya gini kalo mau manggil logging setelah runApp
	//tapi gimana kalo runApp ini ada error?
	fmt.Println(hasil)
}

func main() {
	runApp(0)
}
