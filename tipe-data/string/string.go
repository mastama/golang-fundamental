package main

import "fmt"

/*
* nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.
Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data).
Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya.

	Zero value dari string adalah "" (string kosong).
	Zero value dari bool adalah false .
	Zero value dari tipe numerik non-desimal adalah 0 .
	Zero value dari tipe numerik desimal adalah 0.0 .
*/

/*
* Selain menggunakan tanda quote, deklarasi string juga bisa dengan tanda grave accent/backticks ( ` )
 */
var message = `nama saya adalah "Singgih Pratama"
Yoroseku onegaishimasu
saya adalah developer`

func main() {

	code := 00
	fmt.Printf("%s dengan code %d", message, code)
}
