package main

import "fmt"

// * di go jika kita membuat variabel yang memiliki nilai assign dari variabel lain, maka variabel yang barusan di deklarasikan itu tidak mereferensi variabel lain tersebut
// * sebaliknya dia akan membuat copy variabel dan memori penyimpananya tetaplah berbeda
// * jika kita inin membuat variabel itu reference by variabel lainya, kita bisa melakukanya dengan menggunakan pointer
// * jika nilai dari salah satu pointer berubah, maka pointer lain yang memiliki reference sama juga akan berubah

type Alamat struct {
	Kota, Kabupaten, Kecamatan string
}

func main() {
	var daerah1 Alamat = Alamat{
		Kota:      "Semarang",
		Kabupaten: "Kendal",
		Kecamatan: "Sukorejo",
	}
	// * daerah 2 disini akan me reference ke daerah 1 yang mana jika daerah 2 berubah maka daerah 1 juga akan berubah
	var daerah2 *Alamat = &daerah1

	daerah2.Kecamatan = "Patean"

	fmt.Println(daerah1)
	// fmt.Println(daerah1.Kota)
	fmt.Println(daerah2)
	// fmt.Println(daerah2.Kota)

	fmt.Println("======================")

	// * jika kita mendeklarasikan ulang type nya maka otomatis akan dibuatkan memori baru sehingga bukan lagi reference
	// * jika kita melakukan deklarasi ulang seperti ini maka variabel lain tidak akan berubah walaupun dia merupakan pointer dari al
	var daerah3 *Alamat = &daerah1
	// * ini akan membuat memory baru dan variabel lainya walaupun di deklarasi merupakan reference juga tidak akan berubah
	daerah3 = &Alamat{
		Kota:      "Padang",
		Kabupaten: "Babule",
		Kecamatan: "Walahe",
	}
	fmt.Println(daerah3)

	// * jika kita ingin mendeklarasikan ulang bariabel dan ingin mengubah seluruh variabel yang memiliki reference sama juga ikut berubah
	// * maka kita bisa menggunakan tanda * di awal nama variabel baru itu
	fmt.Println("=============================")
	var daerah4 *Alamat = &daerah1
	*daerah4 = Alamat{
		Kota:      "Papua",
		Kabupaten: "Yahiho",
		Kecamatan: "Yahaha",
	}
	fmt.Println(daerah1)
	fmt.Println(daerah2)
	fmt.Println(daerah3)
	fmt.Println(daerah4)
	fmt.Println("=============================")
}
