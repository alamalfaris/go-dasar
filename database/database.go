package database

import "fmt"

var connection string

func init() { //func/method init ini akan dijalanin ketika package ini dipanggil/diimport
	fmt.Println("func init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
