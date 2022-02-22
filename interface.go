package main

import "fmt"

// interface di golang mirip dengan di java, dia berisi deklarasi method
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// seperti interface pada umumnya dia harus ada implementasinya
// cuman kalo di golang gak perlu bikin class khusus implement kaya java
// dia bisa langsung, tinggal bikin struct dan method-struct nya aja
type Person struct {
	Name string
}

func (person Person) GetName() string { //method ini otomatis ter-connect dengan interface HasName
	return person.Name
}

func main() {
	alam := Person{"alam"}
	SayHello(alam)
}
