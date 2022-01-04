package main

import "fmt"

func main() {
	// * slice adalah tipe data di dalam golang yang mirip sebuah array namun bisa di append
	// * slice memiliki index, panjang dan kapasitas
	// * jika array berubah maka slice juga akan berubah dan begitupun sebaliknya
	// * jika nanti array tidak bisa menampung isi dan di slice masih menambah isinya dengan append atau yang lainya, maka slice akan membuat array baru

	// ? membuat slice dengan membuat array terlebih dahulu
	mounth := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	fmt.Println(mounth)
	newSlice := mounth[10:12]
	fmt.Println(newSlice)
	// * jika kita melakukan perintah ini maka array mount juga akan berubah
	// newSlice[0] = "Bukan November"
	// fmt.Println(newSlice)
	// fmt.Println(mounth)

	appSlice := mounth[6:12]
	fmt.Println(appSlice)
	// * untuk mengetahui panjang slice saat ini
	fmt.Println(len(appSlice))
	// * untuk mengetahui kapasitas maksimal slice
	fmt.Println(cap(appSlice))

	fmt.Println("=========================")

	// ? membuat slice secara langsung
	// * kita juga bisa membuat slice secara langsung tanpa membuat array terlebih dahulu dan hal ini yang lebih di rekomendasikan
	dirSlice := make([]string, 5, 5)
	dirSlice[0] = "Fina"
	dirSlice[1] = "Zahra"
	dirSlice[2] = "Vanessa"
	dirSlice[3] = "Rani"
	dirSlice[4] = "Yolanda"
	// ! kalau ini dilakukan maka akan terjadi error karena akan membuat slice menampung data diluar kapasitasnya
	// dirSlice[5] = "Gina"
	fmt.Println(dirSlice)
	fmt.Println(len(dirSlice))
	fmt.Println(cap(dirSlice))

	fmt.Println("=========================")

	// ? copy slice
	// * jika kita mengubah nilai salah satu atau semua slice hasil copy maka slice parent nya tidak akan berubah
	copySlice := make([]string, len(dirSlice), cap(dirSlice))
	copy(copySlice, dirSlice)
	copySlice[0] = "bukan fina"
	fmt.Println(copySlice)
	fmt.Println(dirSlice)

	fmt.Println("=========================")

	// ? hati hati saat membuat slice dan array
	// ! array
	exArray := [...]int{1, 2, 3, 4, 5, 6}
	// ! slice
	// * ini digunakan untuk membuat slice yang belum diketahui panjang dan kapasitasnya
	exSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(exArray)
	fmt.Println(exSlice)

}
