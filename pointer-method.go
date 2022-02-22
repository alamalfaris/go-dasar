package main

import "fmt"

type Siswa struct {
	Name     string
	Age      int
	Address  string
	Grade    string
	Activity string
}

func (siswa *Siswa) Belajar() { //ditambah * supaya jadi pointer
	siswa.Activity = "Sedang belajar"
}

func Istirahat(siswa *Siswa) { // ditambah * supaya jadi pointer
	siswa.Activity = "Lagi istirahat"
}

// kenapa perlu di pointer?
// supaya hemat memory, nggak duplikat2 tiap masuk ke function

func main() {
	var dina Siswa
	dina.Name = "dina"
	dina.Age = 13
	dina.Address = "Kramat Jati"
	dina.Grade = "2-C"

	// dina.Belajar()
	// fmt.Println(dina)

	// Istirahat(&dina) // ditambah & supaya jadi pointer
	// fmt.Println(dina)

	var doni *Siswa = &dina // ditambah * dan & supaya jadi pointer
	//dono := &dina //bisa juga ditulis gini supaya jadi pointer
	doni.Name = "doni"
	fmt.Println(doni)
	fmt.Println(dina)
}
