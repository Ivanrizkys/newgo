package main

import "fmt"

func main() {
	// * contant adalah varibel yang tidak bisa di rubah type datanya setelah deklarasi
	// * jika kita mendeklarasikan constant tanpa menggunakanya itu tidak apa apa dan berbeda dengan variabel yang akan menyebabkan error
	const nama = "Ivan Rizky Saputra"
	// * pendeklarasian type data disini bersifat opsional
	const namaLengkap string = "Ivan Rizky Saputra"
	fmt.Println(nama)
	fmt.Println(namaLengkap)

	// * deklarasi multipe constant
	const (
		kota      = "Kendal"
		kecamatan = "Sukorejo"
	)
	fmt.Println(kota)
	fmt.Println(kecamatan)
}
