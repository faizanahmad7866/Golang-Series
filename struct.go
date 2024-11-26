package main

import "fmt"

func main() {
	fmt.Println("Struct:")
	type Person struct {
		Name string
		Age  int
	}
	var q Person = Person{"Alice", 30}
	fmt.Println("q =", q)
	fmt.Println("q.Name =", q.Name)
	fmt.Println("q.Age =", q.Age)

	//other example

	type Car struct {
		Name  string
		Model string
		Year  int
	}

	m := Car{"odie", "new", 20224}
	fmt.Println("m =", m)
	fmt.Println("m.Name =", m.Name)
	fmt.Println("m.Model =", m.Model)
	fmt.Println("m.Year =", m.Year)

}
