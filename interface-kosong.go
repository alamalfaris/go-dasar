package main

// interface kosong, digunakan jika kita mau bisa menerima atau return
// semua tipe data
// contoh real penggunaaan interface kosong adalah fmt.Println
// dia bisa print semua tipe data kan
import "fmt"

// atau kita bisa bikin sendiri
func fungsi(x interface{}) interface{} {
	if x == 1 {
		return 1
	} else if x == 2 {
		return true
	} else {
		return "fungsi"
	}
}

func main() {
	// fmt.Println("string") // klo ctrl + klik kanan ntar keliatan dia pake interface kosong
	// fmt.Println(1)
	// fmt.Println(false)
	var data interface{} = fungsi(2)
	fmt.Println(data)
}
