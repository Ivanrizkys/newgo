package main

import "fmt"

func main() {
	// * string dibuat dengan menggunakan cara menuliskanya diantara tanda petik dua
	fmt.Println("Ivan")
	fmt.Println("Ivan Rizky")
	fmt.Println("Ivan Rizky Saputra")

	// * ini merupakan perintah untuk menghitung karakter dalam string
	fmt.Println(len("eko"))
	// * ini adalah perintah untuk mengambil indeks ke berapa yang di tentukan di dalam kurung kotak
	// * namun yang diambil adalan binary code nya
	// * dibawah ini yang diambil adalah binary code dari i
	fmt.Println("Ivan Rizky"[0])
	// * dibawah ini yang diambil adalah binary code dari a
	fmt.Println("Ivan Rizky Saputra"[2])
}
