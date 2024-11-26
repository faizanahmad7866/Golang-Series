//channel common term goroutines communication
//channel are used for communication between goroutines (concurrent threads).They can be synchronous os asynchronous

package main

import "fmt"

func main() {

	fmt.Println("Channel:")
	u := make(chan int) //unbuffered channel of integers
	fmt.Println("u=", u)
}
