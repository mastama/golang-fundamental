package main

import "fmt"

func main() {
	a := 10
	b := 12
	c := 3
	d := b / c

	fmt.Println(a + b/c)
	fmt.Println((a % c) * c)
	fmt.Println(d != c)

}
