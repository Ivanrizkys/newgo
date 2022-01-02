package main

import "fmt"

func main() {
	// * hasil ini akan benar
	const angka32 = 127
	const angka64 = int64(angka32)
	const angka16 = int16(angka32)

	fmt.Println(angka32)
	fmt.Println(angka64)
	fmt.Println(angka16)

	fmt.Println(("=============================="))

	// * hasil ini akan salah
	const bilangan32 = 129
	const bilangan64 = int64(bilangan32)
	// ! ini akan error karena int 8 tidak bisa menampung angka 129
	// const bilangan16 = int8(bilangan32)

	fmt.Println(bilangan32)
	fmt.Println(bilangan64)
	// fmt.Println(bilangan16)

	fmt.Println(("=============================="))

	// * Konversi type data
	var nama = "Ivan"
	// * ini akan menghasilkan byte yang merupakan aliases dari uint8
	var index = nama[0]
	var iString = string(index)
	fmt.Println(nama)
	fmt.Println(iString)
}
