package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now() //dapatkan waktu saat di exec

	fmt.Println(timeNow)
	fmt.Println(timeNow.Year())
	fmt.Println(timeNow.Month())
	fmt.Println(timeNow.Day())
	fmt.Println(timeNow.Hour())
	fmt.Println(timeNow.Minute())
	fmt.Println(timeNow.Second())
	fmt.Println(timeNow.Nanosecond())

	makeTime := time.Date(2022, time.February, 13, 11, 0, 0, 0, time.UTC)
	fmt.Println(makeTime)

	parseTime, _ := time.Parse("2006-01-02", "2022-02-13")
	fmt.Println(parseTime)
}
