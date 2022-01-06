package main

import "fmt"

// * nil hanya bisa digunakan utnuk beberapa type data seperti interface, function, map, slice, pointer, chanel

func getMap(nama string) map[string]string {
	if nama == "" {
		return nil
	}
	return map[string]string{
		"Nama": nama,
	}
}

func main() {
	nama := getMap("ipan")
	if nama == nil {
		fmt.Println("nama Kosong")
	} else {
		fmt.Println(nama["Nama"])
	}
}
