package main

import "fmt"

// * di golang kita bisa menggunakan interface untuk digunakan sebagai contract
// * kita tidak bisa memberikan isi data langusung kepada interface seperti halnya dengan struct
// * untuk itu inteface digunakan dengan menggunakan struct sebagai implementasinya

// * deklarasi interface
type Hasname interface {
	getName() string
}

func sayHello(hasname Hasname) {
	fmt.Println("Hallo", hasname.getName())
}

// * untuk mengakses interface, kita bisa menggunakan struct yang memiliki function getName

type Person struct {
	Name string
}

// * karena struct Person memiliki method getnName() dan return value string maka otomatis masuk ke dalam kontrak Hasname
// * jika nama function dan return value tidak sama dengan yang di deklarasikan interface, maka dia tidak akan dianggap kontrak dengan Hasname
func (person Person) getName() string {
	return person.Name
}

// * contoh lain
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	orang1 := Person{
		Name: "Ivan",
	}
	// * otomatis bisa mengakses function sayHello()
	sayHello(orang1)
	kucing := Animal{
		Name: "Kocheng",
	}
	// * karena struct Animal memiliki method getName dengan return value string, maka otomatis dia terikat kontrak dengan Hasname
	sayHello(kucing)
}
