package main

import (
	"container/list"
	"fmt"
)

//container list disebut juga list, disebut juga double linked list

func main() {
	plant_cd := list.New()
	plant_cd.PushBack("1000")
	plant_cd.PushBack("2000")
	plant_cd.PushBack("3000")
	plant_cd.PushBack("4000")

	//iterate dari depan kebelakang
	// for i := plant_cd.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	//iterate dari belakang kedepan
	for j := plant_cd.Back(); j != nil; j = j.Prev() {
		fmt.Println(j.Value)
	}
}
