package main

import "fmt"

func main()  {
	// arrays
	{ 
	// var a []int
	a:=[3]int{41,42,43}
	fmt.Printf("%#v\n",a)

	b:=[...]string{"hello","my","friends"}
	fmt.Printf("%v",b)

	var c [2]int
	c[0]=1
	c[1]=2
	fmt.Println(c)

	fmt.Printf("a is of type %T\tbut the length is %v\n",a,len(a))
	fmt.Printf("b is of type %T\tbut the length is %v\n",b,len(b))
	fmt.Printf("c is of type %T\tbut the length is %v\n",c,len(c))

	
	}

	

	
}