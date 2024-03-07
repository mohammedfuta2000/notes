package main

import "fmt"

// a func that retruns an int
func foo() int  {
	return 31
}
// a func that returns a differnet func
func bar() (func()int){
	return foo
}


func main() {
	y:=bar()
	fmt.Println(y())

	// THIS IS COO;L;;
	// fmt.Printf("%T",fmt.Println)
	// fmt.Println(x)

	result:=opp(1,1,add)
	fmt.Println(result)
}
func opp(a,b int,f func(int,int)int) int {
	return f(a,b)
}

func add(a,b int) int  {
	return a+b
}
func sub(a,b int )int   {
	return a-b
}