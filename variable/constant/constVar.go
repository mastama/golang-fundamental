package main

import "fmt"

/*
* penggunan const pada pembuatan varible.
* jika menggunaakn const hanya bisa untuk character, string, boolean, numeric dan tidak bisa menggunakan :=
 */

const Pi = 3.14

const (
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

const (
	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
)

// deklrasi multiple konstanta dalam satu baris:
const four, five, six int = 4, 5, 6

func main() {
	fmt.Println(Pi)
	fmt.Println(StatusAccepted)
	fmt.Println(Big, Small, Truth)

	const (
		today string = "saturday"
		sekarang
	)
	isToday := true

	/*
	* - today dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
	* - sekarang dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
	* - isToday2 dideklarasikan dengan metode type inference dengan tipe data bool dan nilainya true
	 */

	fmt.Println(today, sekarang, isToday, four, five, six)
	fmt.Println(today, sekarang, isToday, four+five+six)
}
