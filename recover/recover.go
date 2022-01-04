package main

import "fmt"

// * untuk menangkap error yang dikirimkan oleh panic maka kita bisa melakukanya dengan recover

// * ini adalah function yang akan menghandler jika error terjadi
func errorHandle() {
	message := recover()
	// * kode di dalam if ini hanya akan di eksekusi ketida terjadi error
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai di jalankan")
}
func runApp(err bool) {
	defer errorHandle()
	if err {
		// * mengirimkan error
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
	runApp(true)
	// * jika terjadi error dan kita menhandle nya maka kode dibawah ini akan tetap di eksekusi
	fmt.Println("Ivan")
}
