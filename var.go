package main

import "fmt"

func main() {
	var x int = 10      //cara declare var-1
	var y = 5           //cara declare var-2
	firstName := "alam" //cara declare var-3

	x = 20
	y = 15
	firstName = "didi"

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("firstName = ", firstName)

	//cara declare banyak variable sekaligus

	var (
		usiaAndi = 20
		usiaBudi = 10
		lastName = "alfaris"
	)

	fmt.Println("usia andi = ", usiaAndi)
	fmt.Println("usia budi = ", usiaBudi)
	fmt.Println("full name = ", firstName+" "+lastName)
}
