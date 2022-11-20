package main

import "fmt"

// Untuk memunculkan nilai bool menggunakan fmt.Printf() , bisa gunakan layout format %t .

var value = (((2+6)%3)*4 - 2) / 3
var isEqual = (value != 2)

func main() {
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}
