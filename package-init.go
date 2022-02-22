package main

import (
	_ "belajar-golang-dasar/database" //tanda _ disebut blank identifier, berguna tuk jalanin func init, tanpa panggil func lain
	_ "fmt"
)

func main() {
	// result := database.GetDatabase()
	// fmt.Println("connection:", database.Connection)
}
