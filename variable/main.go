package main

import "fmt"

func main() {

	// Deklarasi varible dengan menyertakan tipe datanya
	//U can declared variable like this:

	var name string = "Dukun Donat"
	var location string = "Jawa Barat"
	fmt.Println(name, location)

	// deklarasi varibale dengan var
	// also can declared variable like this
	var (
		firstName = "Tama"
		lastName  = "Shine"
		age       = 20
	)
	fmt.Println("Hallo perkenalkan saya", firstName, lastName, "Usia saya", age)
	fmt.Println("Hallo perkenalkan saya"+firstName+lastName+"Usia saya", age) // perhatikan perbedaanya

	// deklarasi varible tanpa var dengan menggunakan fungsi fmt.Printf()
	merkLaptop := "Lenovo"
	price := 150000000
	fmt.Printf("Aku membeli laptop dengan merk %s seharga (%d) \n", merkLaptop, price)
	fmt.Println("---------------------------------")

	// or u can do this
	var pekerjaan string = "programmer"
	perusahaan := "sinau academy"
	since := 2022

	fmt.Println("hallo perkenalan saya", firstName, "saya adalah", pekerjaan+" di", perusahaan+" sejak", since)
}
