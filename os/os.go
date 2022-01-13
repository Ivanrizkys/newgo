package main

import (
	"fmt"
	"os"
)

func main() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}
