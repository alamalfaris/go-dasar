package main

import (
	"fmt"
	"os"
)

//package os berisi fungsionalitas untuk mengakses fitur sistem operasi secara independen
//(bisa digunakan di semua OS)
//lengkapnya bisa cek di documentation package os

func main() {
	argument := os.Args //Args untuk menangkap argumen saat program dijalankan
	fmt.Println("argument:", argument)
	fmt.Println("argument:", argument[0])
	fmt.Println("argument:", argument[1])
	fmt.Println("argument:", argument[2])
	fmt.Println("argument:", argument[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("name:", name)
	} else {
		fmt.Println("error:", err.Error())
	}
}
