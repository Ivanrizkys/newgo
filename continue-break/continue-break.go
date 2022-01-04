package main

import "fmt"

func main() {
	// ? continue
	// * Jika ada keyword continue maka program akan di skip ke perulangan selanjutnya dan kode program setelah keyword ini tidak akan dijalankan

	// * perulangan yang hanya mencetak angka negatif
	for number := 0; number <= 11; number++ {
		if number%2 == 0 {
			continue
		}
		fmt.Println(number)
	}

	fmt.Println("===================")

	// ? break
	// * jika terdapat keyword ini maka perulangan akan dihentikan semuanya

	// * jika sampai angka 5 maka perulangan tidak lagi dijalankan
	// * hanya akan mencetak angka sebelum 5
	for angka := 0; angka < 10; angka++ {
		if angka == 5 {
			break
		}
		fmt.Println(angka)
	}
}
