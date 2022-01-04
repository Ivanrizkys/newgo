package main

import "fmt"

// * panic adalah function yang akan menhentikan aplikasi namun tetap akan memanggil defer jika ada

func endApp() {
	fmt.Println("End app")
}

func aplication(err bool) {
	defer endApp()
	if err {
		panic("Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// * aplikasi tidak error
	aplication(false)
	// ! aplikasi error
	// * panic akan ditampilkan setelah defer
	aplication(true)
}
