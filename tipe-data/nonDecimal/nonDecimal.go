package main

import "fmt"

var (
	positiveNumber uint8 = 69
	negativeNumber int64 = -1515468468468
)

var i int = 42
var l float64 = float64(i)
var u uint = uint(l)

func main() {
	/*
	* contoh penggunaan Type Conversion | The expression T(v) converts the value v to the type T.
	 */
	const f = "Ini adalah contoh penggunana nonDecimal data type %T(%v)\n"
	fmt.Printf(f, positiveNumber, positiveNumber)
	fmt.Printf(f, negativeNumber, negativeNumber)

	const j = "%T(%v)\n"
	fmt.Printf(j, i, i)
	fmt.Printf(j, l, l)
	fmt.Printf(j, u, u)
}
