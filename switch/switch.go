package main

import "fmt"

func main() {
	nama := "Saipul"

	switch nama {
	case "Febri":
		fmt.Println("Hallo Febri")
	case "Ipan":
		fmt.Println("Hallo Ipan")
	default:
		fmt.Println("Buang Semangka Buah Kedodondong, Hai cantik kenalan dong")
	}

	// * di switch kita juga bisa melakukan deklarasi secara shorthand
	switch lenNama := len(nama); lenNama > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Sudah benar")
	}

	// * di golang kita juga bisa menggunakan switch tanpa expression
	umur := 20
	switch {
	case umur > 19:
		fmt.Println("Sudah dewasa")
	case umur <= 19:
		fmt.Println("Belum dewasa")
	default:
		fmt.Println("Gatau lagi deh")
	}
}
