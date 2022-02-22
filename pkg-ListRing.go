package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// ring itu seperti list tapi berbentuk lingkaran, artinya gak ada ujungnya
	// utk ring perlu di declare dulu panjangnya
	var data *ring.Ring = ring.New(5)

	// masukin data, pake iterasi
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// ring bisa di println tapi, hasilnya gak keliatan
	// makanya untuk liat value nya bisa pake fungsi bawaan dari ring .Do
	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
