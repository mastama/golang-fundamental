package main

import (
	"fmt"
)

var joni = 9.80
var mean = 8.8

func main() {

	// Template %f digunakan untuk memformat data numerik desimal menjadi string.
	/*
		* Jumlah digit yang muncul bisa dikontrol menggunakan %.nf ,
		tinggal ganti n dengan angka yang diinginkan. Contoh: %.3f
	*/
	fmt.Printf("bilangan desimal: %f\n", joni)
	fmt.Printf("bilangan desimal mean: %.3f", mean)
}
