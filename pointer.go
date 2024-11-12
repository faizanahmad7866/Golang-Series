// a pointer holds the memory adress of a variable For example ,var p *int is a pointer to an integer
package main

import "fmt"

func main() {

	fmt.Println("Pointer:")
	var s int = 10
	var t *int = &s
	fmt.Println("s=", s)
	fmt.Println("t=", t)
	fmt.Println("t points to", *t) // dereferencing the pointer to get the value of the variable


	
	//short way declare variables
	u := 15
	k := &u
	fmt.Println("u=", u)
	fmt.Println("k=", k)
	fmt.Println("k points to", *k) // dereferencing the pointer to get the value of the variable
	*k = 20
	fmt.Println("u after update =", u)
}
