package main

import "fmt"

func main() {
	// ? standard operation
	// + untuk penjumlahan
	// - untuk pengurangan
	// * untuk perkalian
	// / untuk pembagian
	// % untuk sisa pembagian
	var a = 10 + 10
	fmt.Println(a)
	fmt.Printf("Line 1 - Value of a is %d\n", a)

	// ? augmented operation
	var c = 20
	c += 20
	fmt.Println(c)

	// ? unary operation
	// hasilnya akan 41 karena c ditambah 1
	c++
	fmt.Println(c)
}
