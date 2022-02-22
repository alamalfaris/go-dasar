package main

import "fmt"

func main() {
	nama := "farhan"

	switch nama {
	case "farhan":
		fmt.Println("namaku", nama)
	case "alam":
		fmt.Println("namaku", nama)
	default:
		fmt.Println("default")
	}

	//switch dgn short statement
	switch length := len(nama); length > 10 {
	case true:
		fmt.Println("nama kepanjangan")
	case false:
		fmt.Println("nama udah aman")
	}

	//switch tanpa mention variable yg mau di cek
	switch {
	case nama == "farhan":
		fmt.Println("namaku", nama)
	case nama == "alam":
		fmt.Println("namaku", nama)
	default:
		fmt.Println("default")
	}
}
