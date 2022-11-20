package main

import "fmt"

var nilai = 2
var point = 8840.0

func main() {
	if nilai == 10 {
		fmt.Printf("Anda Lulus dengan nilai %d sempurna", nilai)
	} else if nilai > 6 {
		fmt.Printf("Anda Lulus dengan nilai %d", nilai)
	} else {
		fmt.Printf("Anda tidak lulus karena nilai anda %d, dibawah KKM!", nilai)
	}

	println("----------------------------------")

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
