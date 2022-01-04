package main

import "fmt"

// * ini adalah custom type yang bisa di assign di sebuah variabel nantinya
type Anggota struct {
	ketua, wakil, sekretaris string
	kelas                    int
}

func main() {
	// * di dalam go kita bisa menggunakan custom type dengan menggunakan struct
	ipa := Anggota{
		ketua:      "Ivan Rizky Saputra",
		wakil:      "Nanananan",
		sekretaris: "Yana Fara",
		kelas:      10,
	}
	// * untuk mengakses variabel di dalam struct kita bisa menggunakan seperti di bawah ini
	fmt.Println(ipa.ketua)
}
