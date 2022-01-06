package main

import (
	"fmt"
	"reflect"
)

func getString() interface{} {
	return "Hallo"
}

func main() {
	// * type asertion biasa kita gunakan jika kita ingin secara implisit memberikan type kepada variabel yang belum memiliki type
	// * variabel yang belum memiliki type ini biasanya berasal dari return value function dengan interface kosong
	str := getString()
	hasilString := str.(string)
	// ! ini akan menghasilkan panic
	// hasilInt := str.(int)
	fmt.Println(str)
	fmt.Println(hasilString)
	// fmt.Println(hasilInt)
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(hasilString))

	fmt.Println("=============================")
	strr := getString()
	switch value := strr.(type) {
	case string:
		fmt.Println("Bertipe data string, value:", value)
	case int:
		fmt.Println("Bertipe data", value)
	default:
		fmt.Println("Gatau lah males")
	}
}
