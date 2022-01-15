package main

import (
	"fmt"
	"strings"
)

func main() {
	// * akan jika parameter pertama terdapat kata di parameter kedua maka akan mengembalikan nilai true
	fmt.Println(strings.Contains("ivan rizky saputra", "ivan")) // ? true
	// * ini juga bersifat case sensitive
	fmt.Println(strings.Contains("Ivan rizky saputra", "ivan")) // ? false
	fmt.Println(strings.Contains("Ivan rizky saputra", "susi")) // ? false

	fmt.Println("============================")

	// * digunakan untuk men split / memisahkan string menjadi slice
	fmt.Println(strings.Split("ivan rizky saputra", " ")) // ? slice[ivan rizky saputra]

	fmt.Println("============================")

	// * parse semua string ke huruf kecil
	fmt.Println(strings.ToLower("IvAn Rizky SaPutra")) // ? ivan rizky saputra
	fmt.Println(strings.ToUpper("IvAn Rizky SaPutra")) // ? IVAN RIZKY SAPUTRA

	fmt.Println("============================")

	// * bisa digunakan untuk menghapus karakter di awal dan di akhir string
	fmt.Println(strings.Trim("     ivan rizky saputra     ", " "))        // ? ivan rizky saputra
	fmt.Println(strings.Trim("     ivan rizky saputra", " "))             // ? ivan rizky saputra
	fmt.Println(strings.Trim("llllllivan rizky saputralllllllllll", "l")) // ? ivan rizky saputra

	fmt.Println("============================")

	// * digunakan untuk me replace string
	// * hanya  satu yang di replace
	// * parameter terakhir adalah berapa string yang akan di replace
	fmt.Println(strings.Replace("ivan ivan ivan yahya", "ivan", "galih", 2)) // ? galih galih ivan yahya
	// * akan me replace semua
	fmt.Println(strings.ReplaceAll("ivan ivan ivan yahya", "ivan", "galih")) // ? galih galih galih yahya
}
