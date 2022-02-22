package main

import (
	"errors"
	"fmt"
)

//sudah ada bawaan interface untuk error
//ada di package import "errors"

//udah biasa di golang, kalo function nya bisa ada error
//maka function nya return 2 value, 1 hasilnya, 1 errornya
func Pembagian(val1 int, val2 int) (int, error) {
	if val1 == 0 || val2 == 0 {
		return 0, errors.New("gak bisa bagi kalo ada 0 nya")
	} else {
		result := val1 / val2
		return result, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err)
	}
}
