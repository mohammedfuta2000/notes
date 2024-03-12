package main

import "fmt"

func main() {

	c := make(chan int)

	go func(chan<- int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}(c)

	// u could use a wait group. 
	// but running bar in the main goroutine blocks it enough
	for v := range c{
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
