package main

import "fmt"

func main() {
	f, i := incrementor()
	fmt.Println(f())
	f()
	f()
	f()
	fmt.Println()
	fmt.Println(*i)
}

func incrementor() (func() int, *int) {
	x := 0
	return func() int {
		x++
		return x
	}, &x
}
