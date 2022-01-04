package main

import (
	"fmt"
	"reflect"
)

// ? function dengan parameter
func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hallo", firstname, lastname)
}

// ? function tanpa parameter
func sayHello() {
	fmt.Println("Hallo Ivan")
}

// ? function dengan return value
func helloName(nama string) string {
	return "Hallo " + nama
}

// ? function dengan multiple return value
// * di golang kita bisa membuat function me return value sesuai dengan yang kita tentukan
func getFullname(firstname string, middlename string, lastname string) (string, string, string) {
	return firstname, middlename, lastname
}

// * di golang kita bisa melakukan deklarasi return variabel pada tempat pendeklarasian type return
// * jika return value semuanya bertype sama maka bisa hanya di deklarasikan di akhir
func family() (bapak, ibu, anak string) {
	bapak = "Handra"
	anak = "Fikri"
	ibu = "Doni"
	// * kita tidak perlu lagi mendeklarasikan apa yang di return karena sudah di definisikan diatas
	return
}

// ? variadic arguments in function
// * kita juga bisa membuat arguments yang bersifar dynamic di go
// * caranya adalah dengah kita menambahkan variadic arguments di dalam function
// * variadic ini haruslah berada di akhir arguments dan tidak boleh diawal ataupun di tengah
func sumAll(number ...int) int {
	fmt.Println(reflect.TypeOf(number))
	fmt.Println(number)
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

// ? function as parameter
// * di golang kita bisa membuat sebuah function dapat menerima parameter dengan type function
// function params
func filterName(name string) string {
	if name == "Anjing" || name == "anjing" {
		return "..."
	}
	return name
}

// function yang menerima parameter function
func sayName(name string, filter func(string) string) string {
	return "Hai " + filter(name)
}

// * jika parameter nya terlalu panjang maka kita bisa membuatnya dengan menggunakan type declaration
type Filter func(string) string

func sayName2(name string, filter Filter) string {
	return "Hai " + filter(name)
}

// ? anonymus function
type Blocked func(string) bool

func filterBlock(name string, blocked Blocked) string {
	filter := blocked(name)
	if filter {
		return "You are blocked"
	}
	return name
}

func main() {
	sayHello()
	fmt.Println("=======================")
	sayHelloTo("Ivan", "Rizky")
	fmt.Println("=======================")
	hello := helloName("Ivan Rizky Saputra")
	fmt.Println(hello)
	fmt.Println("=======================")
	firstname, middlename, lastname := getFullname("Ivan", "Rizky", "Saputra")
	// * jika kita tidak ingin mengambil seluruh return value nya maka bisa menggunakan _
	// * kode dibawah hanya akan mengambil middlename
	// _, middlename,_ := getFullname("Ivan", "Rizky", "Saputra")
	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
	fmt.Println("=======================")
	a, b, c := family()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("=======================")
	angka := []int{1, 2, 3, 4, 5, 5}
	sum := sumAll(12, 12, 12, 12, 12)
	sum1 := sumAll(angka...)
	fmt.Println(sum)
	fmt.Println(sum1)

	fmt.Println("=======================")
	// * function juga bisa di assign di dalam sebuah variabel dengan cara seperti ini
	hai := helloName
	fmt.Println(hai("Ipan"))

	fmt.Println("=======================")
	// ? function as parameter
	sapa := sayName("anjing", filterName)
	fmt.Println(sayName2("Aisyah", filterName))
	fmt.Println(sapa)

	fmt.Println("=======================")
	// ? anonymous function
	// * kita bisa membuat function anonymous di dalam sebuah variabel
	fil := func(name string) bool {
		return name == "Ivan"
	}
	fmt.Println(filterBlock("Ivan", fil))
	// * membuat function anonymous langsung di dalam parameter
	fil2 := filterBlock("Hendra", func(name string) bool {
		return name == "Ivan"
	})
	fmt.Println(fil2)

}
