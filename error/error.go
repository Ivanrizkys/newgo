package main

import (
	"errors"
	"fmt"
)

// * function ini akan memberikan 2 return value
func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return nilai, errors.New("Pembagian tidak boleh 0")
	}
	return nilai / pembagi, nil
}

func main() {
	// * cara membuat error
	err := errors.New("Ini error")
	fmt.Println(err.Error())

	// * cara memanggilnya kita bisa melakukan mirip mirip dengan destructuring
	nilai, err := pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasilnya:", nilai)
	} else {
		fmt.Println("Error message:", err.Error())
	}
}
