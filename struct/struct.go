package main

import "fmt"

// * ini adalah custom type yang bisa di assign di sebuah variabel nantinya
type Anggota struct {
	ketua, wakil, sekretaris string
	kelas                    int
}

// * di golang kita juga bisa membuat struct seolah olah memiliki sebuah function / mehotd
// * parameter anggota tadi bisa dinamai dengan yang lainya misalkan a, b, dll
func (anggota Anggota) sayHelloFromKetu(name string) {
	fmt.Println("Hallo", name, "nama saya", anggota.ketua)
}

// * contoh lain
func (a Anggota) sayHuuuu() {
	fmt.Println("Huuuuuu form", a.ketua)
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

	fmt.Println("==================")

	// * nah disini kita bisa mengakses sayHelloFromKetu dengan menggunakan varuabel yang sudah di assign menggunakan struct tadi
	ipa.sayHelloFromKetu("ivan")
	ipa.sayHuuuu()
}
