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

func main() {
	fmt.Println(Pi)
	fmt.Println(StatusAccepted)
	fmt.Println(Big, Small, Truth)
}
