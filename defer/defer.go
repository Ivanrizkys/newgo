package main

import "fmt"

func logging() {
	message := recover()
	// * ini akan menangkap error agar aplikasi tetap berajalan
	if message != nil {
		fmt.Println("Aplikasi error dengan message:", message)
	}
	fmt.Println("Selesai Memanggill Function")
}

func runApplication(number int) {
	// * deffer function akan dijalankan terakhir setelah function runApplication di jalankan
	// * walaupun dalam runApplication terdapat error maka defer function akan tetap di jalankan
	defer logging()
	fmt.Println("Run Application")
	// ! perintah ini akan menghasilkan error
	value := 10 / number
	fmt.Println(value)
}

func main() {
	runApplication(10)
	runApplication(0)
}
