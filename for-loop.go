package main

import "fmt"

func main() {

	// di golang looping cuma ada for aja
	counter := 1

	for counter <= 10 {
		fmt.Println("counter ke:", counter)
		counter++
	}

	//bisa kaya gini juga
	for i := 1; i <= 10; i++ {
		fmt.Println("counter ke:", i)
	}

	//for range, looping untuk akses data collection
	names := []string{"alam", "adi", "bagus", "cecil", "danny"}
	for i, name := range names {
		fmt.Println("index", i, "=", name)
		//fmt.Println(name) //kalo mau print name nya aja, i nya diganti _ (underscore)
	}

	//kita cobain looping map
	students := make(map[string]string)
	students["name"] = "alam"
	students["address"] = "jakarta"

	for key, value := range students {
		fmt.Println(key, "=", value)
	}
}
