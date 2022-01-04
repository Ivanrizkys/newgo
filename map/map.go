package main

import "fmt"

func main() {
	// * tiipe data map ini adalah tipe data yang memiliki key dan value
	// * untuk mengakses value kita harus menggunakan key nya, dan oleh karena itu key nya haruslah unix
	// * jika kita mendeklarasikan lagi key yang sudah ada maka key lama akan di replace begitupun dengan value nya
	// ? membuat map
	family := map[string]string{
		"ibu":   "Fika Amanda",
		"bapak": "Reza Putra",
		"anak":  "Dea Hanissa",
	}

	// * jika kita mendeklarasikan key dan value yang mana key nya belum ada di dalam variabel maka itu otomatis ditambahkan
	family["nenek"] = "Nisa"

	// * untuk men delete elelement di dalam map
	delete(family, "anak")

	// ? membuat map dengan deklarasi tanpa isi
	organisasi := make(map[string]int)
	organisasi["ketua"] = 1
	fmt.Println(organisasi)

	fmt.Println(family)
	fmt.Println(family["ibu"])
}
