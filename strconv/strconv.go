package main

import (
	"fmt"
	"strconv"
)

func main() {
	// * mengubah string menjadi boolean
	// * return hasil dan error
	bolean, _ := strconv.ParseBool("false")
	fmt.Println(bolean)

	// * mengubah string menjadi integer
	// * parameter kedua adalah base nya bisa hexa, oktal, dan sebagainya
	integer, _ := strconv.ParseInt("90", 10, 64)
	fmt.Println(integer)

	// * mengubah boolean ke string
	strBol := strconv.FormatBool(true)
	fmt.Println(strBol)

	// * mengubah integer menjadi string
	strInt := strconv.FormatInt(200, 10)
	fmt.Println(strInt)

	// * mengubah string ke integer
	integer1, _ := strconv.Atoi("2000")
	fmt.Println(integer1)

	// * mengubah integer ke string
	stringg := strconv.Itoa(3000)
	fmt.Println(stringg)
}
