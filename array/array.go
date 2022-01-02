package main

import "fmt"

func main() {
	// * tipe data di array haruslah sama
	// * jumlahnya pun tidak boleh lebih dari yang di deklarasikan
	// * tidak bisa bertambah setelah dibuat jika kapasitasnya sudah di deklarasikan

	// ? deklarasi secara manual
	var siswa [5]string
	siswa[0] = "Hendra"
	siswa[1] = "Fikri"
	siswa[2] = "Sueb"
	siswa[3] = "Hana"
	siswa[4] = "Rosi"
	// ! ini akan error karena diluar dari yang di deklarasikan
	// siswa[5] = "Hendra"

	fmt.Println(siswa[0])
	fmt.Println(siswa[1])
	fmt.Println(siswa[2])
	fmt.Println(siswa[3])
	fmt.Println(siswa[4])
	fmt.Println(siswa)

	fmt.Println("===========================")
	var any = [...]string{"ivan", "rizky", "saputra"}
	fmt.Println(any)

	fmt.Println("===========================")

	// ? deklarasi secara langsung
	var nilai = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nilai)

	// ? mengakses panjang dari array
	// * ini adalah cara untk mengetahui panjang dari array dan bukan jumlah data di dalam array
	fmt.Println(len(siswa))
	fmt.Println(len(nilai))

	// * index array yang belum kita berikan nilai akan bernilai 0
	var test = [2]int{1}
	fmt.Println(test)
}
