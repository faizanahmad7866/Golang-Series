// map are used to store data in key-value pairs.For example map [string]int is a map with string key and integer value

package main

import "fmt"

func main() {

	fmt.Println("map:")
	r := map[string]int{"Alice": 30, "Bob": 25} // map with string keys and int vlaue
	fmt.Println("r =", r)
	fmt.Println("Alice=", r["Alice"])
	fmt.Println("Bob=", r["Bob"])
	

}
