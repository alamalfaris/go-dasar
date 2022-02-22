package main

import "fmt"

//closure itu fitur sebuah fungsi bisa ubah data di atasnya
func main() {
	nama := "alam" //ini kan awalnya alam

	increment := func() {
		fmt.Println("increment")
		nama = "odi" //nah disini bakal ke ganti, jadi odi
		//makanya harus hati2, jangan sampe salah, takutnya value yg atas malah keganti
	}

	increment()
	fmt.Println(nama)
}
