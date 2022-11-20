package main

import "fmt"

func main() {
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	// Template \t digunakan untuk menambahkan indent tabulasi. Biasa dimanfaatkan untuk merapikan tampilan output pada console.
}

/*
*	Berikut penjelasan statemen operator logika pada kode di atas.
	- leftAndRight bernilai false , karena hasil dari false dan true adalah false .
	- leftOrRight bernilai true , karena hasil dari false atau true adalah true .
	- leftReverse bernilai true , karena negasi (atau lawan dari) false adalah true .
*/
