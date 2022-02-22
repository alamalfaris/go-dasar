package main

import (
	"fmt"
	"reflect"
)

type Orang struct {
	Name   string `required:"true"` // di struct, kita bisa kasih tag di tiap field
	Alamat string
	NIP    string `required:"true"`
}

// dengan tag di field struct, kita bisa bikin common/library untuk validasi
// contonhnya kalo ada Orang namanya gak diisi, trus kita call function ini
// dia akan return false, kalo namanya diisi return nya true
func isValid(data interface{}) string {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return "field " + field.Name + " belum diisi"
			}
		}
	}
	return "SUCCESS"
}

func main() {
	org1 := Orang{"ardi", "Jakarta", "L001"}

	org1Type := reflect.TypeOf(org1)

	//mau cek di struct orang ada berapa field
	fmt.Println(org1Type.NumField())

	//mau cek di struct orang ada nama field apa
	fmt.Println(org1Type.Field(0).Name)

	//mau ambil value dari tag required di struct orang
	fmt.Println(org1Type.Field(0).Tag.Get("required"))

	//coba call isValidName
	org2 := Orang{"", "Jakarta", ""}
	fmt.Println(isValid(org2))
}
