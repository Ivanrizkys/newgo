package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ? Pertama dalam golang bisa membuat variabel dengan mendeklarasikanya terlebih dahulu
	// ? namun pada saat ini haruslan mendeklarasikan type variabel nya juga
	var nama string
	nama = "Ivan Rizky Saputra"
	fmt.Println(nama)

	// ? kita juga bisa langsung meng assign nilai variabel
	// ? jika cara ini dilakukan kita tidak harus mendeklarasikan jenis variabel nya
	var adik = "Aisyah Nisrina Habibah"
	fmt.Println(adik)

	fmt.Println("===================")

	// * namun beda halnya jika kita langsung meng assign variabel yang memilikir type data number
	// * variabel nilaiIvan dan nilaiHendra akan memiliki type int yang memiliki minimal int 32
	var nilaiIvan = 60
	var nilaiHendra = 50
	var nilaiAis = 9.5

	fmt.Println(nilaiIvan)
	fmt.Println(nilaiHendra)
	fmt.Println(nilaiAis)
	// * reflect digunakan untuk mengecek type data dari sebuah varibel
	fmt.Println(reflect.TypeOf(nilaiIvan))
	fmt.Println(reflect.TypeOf(nilaiHendra))
	fmt.Println(reflect.TypeOf(nilaiAis))

	fmt.Println("===================")

	// * oleh karena itu jika kita ingin mendefinisikan int atau float secara expicit maka harus melakukan deklarasi terlebih dahulu dengan type nya
	var nilaiAbduh int8
	nilaiAbduh = 100
	fmt.Println(nilaiAbduh)
	var nilaiSiti float32
	nilaiSiti = 9.3
	fmt.Println(nilaiSiti)

	fmt.Println(reflect.TypeOf(nilaiAbduh))
	fmt.Println(reflect.TypeOf(nilaiSiti))

	fmt.Println("===================")

	// ? deklarasi variabel tanpa var
	// * di golang jika kita ingin mendeklarasikan variabel tanpa var maka harus menggunakan := pada saat inisialisasi namun tidak perlu saat kita melakukan assigment
	negara := "Indonesia"
	fmt.Println(negara)
	// * tidak perlu menggunakan :=
	negara = "Brazil"
	fmt.Println(negara)

	fmt.Println("===================")

	// ? deklarasi multiple variabel sekaligus
	// * cara ini digunakan jika ingin mendeklarasikan beberapa variabel sekaligus
	var (
		sma = "SMA Negri 1 Sukorejo"
		smp = "SMP Negri 1 Sukorejo"
	)
	fmt.Println(sma)
	fmt.Println(smp)

	var (
		angka1 int16   = 100
		angka2 float32 = 3.5
	)
	fmt.Println(angka1)
	fmt.Println(angka2)
}
