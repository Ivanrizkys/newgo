package main

import (
	"fmt"

	// ? blank identifier
	// * agar tidak terjadi error karena mengimport package dan tidak mengggunakanya
	_ "github.com/Ivanrizkys/newgo/database"
)

func main() {
	// * memanggil function dari package database
	// * jika kita tidak akan menanggil function di dalam package maka bisa melakukan dengan blank identifier pada saat import
	// dbName := database.GetDatabase()
	// fmt.Println(dbName)
	fmt.Println("succes")
}
