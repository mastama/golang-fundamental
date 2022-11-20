package main

import (
	"fmt"
)

var stars = 1

func main() {
	switch stars {
	case 5:
		fmt.Println("Congrats you got 4 stars!")
	case 4:
		fmt.Println("Its okay, you still get 4 stars!")
	case 3:
		fmt.Println("You need improvement, cause you got 3 stars!")
	case 2:
		fmt.Println("Danger, you get 2 stars!")
	default:
		fmt.Println("You will get suspend for a while day!"+" your stars only ", stars)
	}

	println("-------------------------------------------")

	var points = 6
	switch points {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	var point = 5
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 5):
		fmt.Println("awesome")
	case (point > 1) && (point < 4):
		fmt.Println("not good")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	println("-------------------------------------------")

	/*
		* Ketika sebuah case terpenuhi, pengecekan
		kondisi tidak akan diteruskan ke case-case setelahnya.
		Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan
		ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya, jadi
		satu case di pengecekan selanjutnya tersebut selalu dianggap benar (meskipun
		aslinya adalah salah). Dalam sebuah switch lebih dari satu fallthrough bisa di
		tempatkan untuk memaksa melanjutkan proses pengecekan ke satu case
		setelahnya.
	*/

	var pointo = 6

	switch {
	case pointo == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case pointo < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println("-------------------------------------")

	/*
		setelah pengecekan case (point < 8) && (point > 3) selesai, akan dilanjut ke
		pengecekan case point < 5 , karena ada fallthrough di situ.
	*/

	/*
		*Seleksi kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi
		kondisi, yang mungkin juga berada dalam seleksi kondisi, dan seterusnya.
		Seleksi kondisi bersarang bisa dilakukan pada if - else , switch , ataupun
		kombinasi keduanya.
	*/

	var value = 7

	if value >= 7 {
		switch value {
		case 10:
			fmt.Println("Nice, you get PERFECT!")
		default:
			fmt.Println("Nice crot!!!")
		}
	} else {
		if value > 5 && value <= 6 {
			fmt.Println("You need more learning!")
		} else {
			fmt.Println("you must hard work!")
			if value == 0 {
				fmt.Println("you're failed! try harder!")
			}
		}
	}
}
