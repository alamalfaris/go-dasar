package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (val UserSlice) Len() int {
	return len(val)
}

func (val UserSlice) Less(i, j int) bool {
	// ceritanya mau urutin based on Age
	return val[i].Age < val[j].Age
}

func (val UserSlice) Swap(i, j int) {
	val[i], val[j] = val[j], val[i]
}

func main() {
	listUsers := UserSlice{
		{"Alam", 24},
		{"Doni", 34},
		{"Akbar", 22},
		{"Dani", 26},
	}

	fmt.Println(listUsers) // before sort
	sort.Sort(listUsers)
	fmt.Println(listUsers) //after sort
}
