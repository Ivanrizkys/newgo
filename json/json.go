package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama  string
	Kelas string
}

func main() {
	// * array of object
	arrObj := []User{{"Ivan", "kelas"}, {"sagiri", "gws"}}
	var jsonData, _ = json.Marshal(arrObj)
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	// * json

}
