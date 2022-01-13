package main

import (
	"flag"
	"fmt"
)

func main() {
	// * package flag bisa digunakan untuk menangkap value yang dikirimkan dari perintah yang kita ketikkan di terminal
	// * ketikkan perintah dibawah ini di terminal
	// ! go run flag.go -host=localhost -user=ipan password=sakayusaka65

	// ? hasil dari flag.String() adalah variabel pointer
	// * parameter pertama adalah key dari flag
	// * parameter kedua adalah default value jika key dari flag tidak memiliki value
	// * parameter ketiga adalah helper
	var host *string = flag.String("host", "localhost", "hostname from database")
	var username *string = flag.String("user", "root", "user from database")
	var password *string = flag.String("password", "sahaja", "password from database")

	// * untuk mem parse seluruh flag agar bisa dipanggil
	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *username)
	fmt.Println("Password:", *password)
}
