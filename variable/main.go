package main

import "fmt"

func main() {

	// Deklarasi varible dengan menyertakan tipe datanya
	//U can declared variable like this:
	// menggunakan konsep manifest typing (Ketika deklarasi, tipe data yg digunakan harus dituliskan juga)

	var name string = "Dukun Donat"
	var location string = "Jawa Barat"
	fmt.Println(name, location)
	fmt.Println("---------------------------------")

	// deklarasi varibale dengan var | penggunaan multiple varible
	// also can declared variable like this
	var (
		firstName = "Tama"
		lastName  = "Shine"
		age       = 20
	)
	fmt.Println("Hallo perkenalkan saya", firstName, lastName, "Usia saya", age)
	fmt.Println("Hallo perkenalkan saya"+firstName+lastName+"Usia saya", age) // perhatikan perbedaanya

	var hero, power, skills = "Deku", 7000, "jalan didinding"
	fmt.Println(hero, "adalah hero yang memiliki power", power, "dan dapat", skills)
	fmt.Println("---------------------------------")

	// deklarasi varible tanpa var dengan menggunakan fungsi fmt.Printf()
	// menggunakan konsep type inference (type data tidak perlu dituliskan)
	merkLaptop := "Lenovo"
	price := 150000000
	fmt.Printf("Aku membeli laptop dengan merk %s seharga (%d) \n", merkLaptop, price)
	fmt.Println("---------------------------------")

	// or u can do this
	var pekerjaan string = "programmer"
	perusahaan := "sinau academy"
	since := 2022

	fmt.Println("hallo perkenalan saya", firstName, "saya adalah", pekerjaan+" di", perusahaan+" sejak", since)
	fmt.Println("---------------------------------")

	// membuat varible kosong atau underscore _ // manis
	var _ = "Belajar Golang Fundamental"
	_ = "ini tidak digunakan"
	bahasa, _ := "Golang", 2019
	fmt.Println("penggunaan _ pada pembuatan varible dikarenakan golang tidak bisa menerima jika ada varible yang unussed di", bahasa)
	fmt.Println("---------------------------------")

	// membuat varible dengan new (pointer string)
	// dalam penggunaakn pointer string harus menggunakan dereference dengan menambahkan tanda asterisk(*) agar yang tampil valuenya
	switching := new(string)
	fmt.Println("ini adalah service switching", switching) //akan menampilkan alamat memori value tsb dalam bentuk notasi heksadesimal
	fmt.Println("ini adalah service dari switching", *switching, "yang dibuat menggunakan dereference dengan menambahkan tanda asterisk(*)")

}
