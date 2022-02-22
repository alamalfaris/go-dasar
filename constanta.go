package main

import "fmt"

//const itu gak bisa di ubah value nya, kecuali hardcode nya diganti
//const itu gak error kalo gak terpakai

func main() {
	const pathHome string = "/home" //cara declare const-1
	const pathProfile = "/profile"  //cara declare const-2

	const ( //cara declare const langsung banyak
		pathMaster = "/master"
		id         = 2001
	)

	fmt.Println(pathHome)
	fmt.Println(pathProfile)
	fmt.Println(id)
}
