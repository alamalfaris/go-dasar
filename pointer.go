package main

import "fmt"

//di golang pake konsep pass by value
//artinya saat variable masuk ke function/method atau ke variable lain
//itu value dari variable itu di duplikat

type Address struct {
	City     string
	Province string
	Country  string
}

func ChangeCountryToIndo(address *Address) { //dikasi * tanda Address nya pointer
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Bekasi", "Jawa Barat", ""}
	//ChangeCountryToIndo(address)
	ChangeCountryToIndo(&address) //dikasi & karna address nya maunya yg pointer
	fmt.Println(address)
	//----------------------------------------------||--------------------------------------------
	// var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// var address5 *Address = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} //cara bikin pointer kalo variable baru
	// address2 := address1                                                    // kalo assign nya kyk gini, value address1 di copy ke address2
	// address2.City = "Bekasi"                                                // kalo di java, kalo value ini diganti, harusnya yang di address1 juga keganti
	// address2.Province = "Jawa Barat"

	// // fmt.Println(address1)
	// // fmt.Println(address2)

	// address4 := &address1
	// address3 := &address1                                    //kalo assign nya kyk gini, address3 me-refer ke address1
	// *address3 = Address{"Malang", "Jawa Timur", "Indonesia"} // kalo pake * trus declare baru gini, nanti semua address akan me-refer address3

	// fmt.Println(address1)
	// fmt.Println(address3)
	// fmt.Println(address4)
	// fmt.Println(address5)

	// //bikin pointer dengan value kosong
	// //keyword nya new
	// var address6 *Address = new(Address)
	// //kita coba isi field nya
	// address6.City = "Surabaya"
	// fmt.Println(address6)
}
