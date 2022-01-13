package database

import "fmt"

var database string

// * function ini akan di jalankan saat ada function di dalam package ini yang di akses
func init() {
	// * test apakan function di jalankan
	fmt.Println("function init di dijalankan")
	database = "Mysql"
}

func GetDatabase() string {
	return database
}
