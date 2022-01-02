package main

import "fmt"

func main() {
	// ? type declaration
	// * digunakan untuk membuat aliases kepada type data yang ada
	// * misalkan kita akan membuat alias maried untuk type data bool

	type Married bool
	var menikah Married = true
	fmt.Println(menikah)
}
