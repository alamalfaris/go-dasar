package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nilai32 int32 = 20
	var nilai64 int64 = int64(nilai32)
	fmt.Println("nilai = ", nilai64)

	nilai := 21
	nilaiStr := strconv.Itoa(nilai) //convert nilai to string
	fmt.Println("nilai = ", nilaiStr)

}
