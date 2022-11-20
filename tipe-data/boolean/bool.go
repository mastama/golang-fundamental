package main

import "fmt"

func main() {

	/*
		*Tipe data bool berisikan hanya 2 variansi nilai, true dan false . Tipe data ini
		biasa dimanfaatkan dalam seleksi kondisi dan perulangan

		noted: Gunakan %t untuk memformat data bool
	*/
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	notExist := false
	fmt.Printf("notExist %t\n", notExist)
}
