package main

import (
	"fmt"
	"regexp"
)

func main() {
	//ceritanya bikin regex utk cek string depannya 'e' belakangnya 'o', tengahnya a-z
	regex := regexp.MustCompile("e([a-z])o")

	//untuk cek apakah string sesuai dengan regex
	fmt.Println(regex.MatchString("eno"))
	fmt.Println(regex.MatchString("ePo"))

	//mau cari dan ambil string yg sesuai regex dari kumpulan kata
	result := regex.FindAllString("eko ako eno ena ezo", -1) //-1 untuk ambil semuanya
	fmt.Println(result)
}
