package main

import "fmt"

func main() {
	nama := "Ivan"
	umur := 19

	if nama == "Ivan" {
		fmt.Println("Hallo Ivan")
	} else if nama == "Winda" && umur == 19 {
		fmt.Println("Hallo winda umur 19")
	} else {
		fmt.Println("Hallo man teman")
	}

	// * di go kita juga bisa mendeklarasikan sesuatu sebelum expression if di jalankan
	if length := len(nama); length > 5 {
		fmt.Println("Nama lebih panjang dari 5")
	} else {
		fmt.Println("Nama Lebih Kecil dari 5")
	}
}
