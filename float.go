package main

import "fmt"

func main() {

	fmt.Println("floa32 and float64:")
	var h float32 = 3.14 //single precision
	fmt.Println("h=", h)

	fmt.Println("Complex numbers:")
	var j complex64 = 2 + 3i  //single precision 4 byte of data
	var k complex128 = 2 + 3i //double precision 8 byte of data
	fmt.Println("j=", j)
	fmt.Println("k=", k)
}
