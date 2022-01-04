package main

import "fmt"

func main() {
	number := 0
	for number <= 10 {
		fmt.Println("Angka saat ini adalah:", number)
		number++
	}

	fmt.Println("===================")

	// * di golang kita bisa menambahkan 2 steatment sebelum kondisi dan sesudah kondisi
	// * steatment di awal hanya akan dijalankan sesaat sebelum for dijalankan
	// * steatment akhir akan selalu dijalankan setelah for dijalankan selama kondisi masih memenuhi
	for angka := 0; angka <= 11; angka++ {
		// * steatment initialisasi angka hanya akan dijalankan sekali
		// * steatment angka ++ akan dijalankan pada akhir for selama kondisi masih memenuhi
		fmt.Println("Angka:", angka)
	}

	fmt.Println("===================")

	// ? for untuk iterasi array, slice, map
	// * for range bisa melakukan iterasi untuk array, slice dan map
	cewe := []string{"Irvana", "Sania", "Tasya", "Mela"}
	for i, nama := range cewe {
		fmt.Println("Cewe ke", i, "namanya", nama)
	}
	// * jika kita tidak ingin menggnakan variabe i maka bisa melakukan seperti di bawah ini
	for _, nama := range cewe {
		fmt.Println(nama)
	}
	fmt.Println("===================")

	// ? looping di map
	kelas := map[string]string{
		"guru":  "Vivi",
		"Ketua": "Galih",
		"Wakil": "Raka",
	}
	for key, value := range kelas {
		fmt.Println(key, "=", value)
	}
}
