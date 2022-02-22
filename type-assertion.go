package main

import "fmt"

//type assertion tu merubah tipe data dari interface kosong

func Asal() interface{} {
	return "fungsi asal"
}

func main() {
	var result interface{} = Asal()
	//var resultString string = result.(string) //meng-convert result jadi string
	//fmt.Println(resultString)

	//TAPI,, kalo sampai salah convert, akan muncul panic
	//program will stopped
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	//handling panic nya bisa pake switch
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}
}
