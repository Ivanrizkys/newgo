package main

import "fmt"

func main() {
	const kalimat1 = "budi"
	const kalimat2 = "Edi"
	const kalimat3 = "Edi"
	// * ini digunakan untuk menggabungkan string
	fmt.Println(kalimat1 + kalimat2)
	// * digunakan untuk membandingkan string apakah sama atau tidak
	// * ini juga bersifat case sensiitive
	fmt.Println(kalimat1 == kalimat2)
	fmt.Println(kalimat2 == kalimat3)

	fmt.Println("=============================")

	const bilangan1 = 20
	const bilangan2 = 40
	fmt.Println(bilangan1 == bilangan2)
	fmt.Println(bilangan1 > bilangan2)
	fmt.Println(bilangan1 < bilangan2)
	fmt.Println(bilangan1 >= bilangan2)
	fmt.Println(bilangan1 <= bilangan2)
	fmt.Println(bilangan1 != bilangan2)
}
