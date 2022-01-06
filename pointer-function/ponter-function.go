package main

import "fmt"

type Addres struct {
	Kota      string
	Kabupaten string
	Kecamatan string
}

// * jika kita ingin membuat struct method usahakan untuk membuatnya dalam bentuk pointer karena ini bisa menghemat memori
func (alamat *Addres) changeKecamatan() {
	alamat.Kecamatan = "Kecamatan " + alamat.Kecamatan
}

func chageKabupaten(alamat Addres) {
	alamat.Kabupaten = "Semarang"
}
func chageKabupatenPointer(alamat *Addres) {
	alamat.Kabupaten = "Semarang"
}

func main() {
	// * kita bisa memberikan pointre sebagai nilai dari parameter sebuah function
	// * hal ini dilakukan jika kita membuat data di dalam struct yang banyak
	// * ketimbang memsaukkanya ke dalam value, kita bisa memberikanya sebagai pointer agar hemat penggunaan memori
	var alamat1 Addres = Addres{
		Kota:      "Bandung",
		Kabupaten: "Malang",
		Kecamatan: "Bali",
	}
	chageKabupaten(alamat1)
	// * jika kita melakukan seperti ini maka data kabupaten di alamat 1 tidak akan berubah
	fmt.Println(alamat1)
	// * namun jika memberikan argument dengan sebuah pointer maka kota di alamat1 akan berubah
	chageKabupatenPointer(&alamat1)
	fmt.Println(alamat1)
	alamat1.changeKecamatan()
	fmt.Println(alamat1)
}
