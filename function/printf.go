package main

import "fmt"

func main() {
	name := "Caprica-Six"
	aka := 6
	fmt.Printf("%s is also known as %d", name, &aka)
}
