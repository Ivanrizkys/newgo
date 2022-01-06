package main

import "fmt"

// * jika kita ingin memberikan type yang bersifat dynamic / di golang, kita bisa menggunakan interface kosong
// * function ini memiliki return value yang dinamis
func acak(angka int) interface{} {
	// * dia bisa be return value angka, string, maupun
	if angka == 1 {
		return 1
	} else if angka == 2 {
		return true
	} else {
		return "hallo"
	}
}

// * kita juga bisa membuat parameter di dalam function nya dinamis dan juga bisa menggunakan variadic function
func hallo(nama ...interface{}) interface{} {
	for _, nama := range nama {
		fmt.Println(nama)
	}
	return nama
}

func main() {
	fmt.Println(acak(1))
	fmt.Println(acak(2))
	fmt.Println(acak(3))

	hallo("ivan", "Yahya")

}
