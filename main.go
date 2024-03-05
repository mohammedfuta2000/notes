package main

import "fmt"

func main() {
	// arrays
	// {
	// 	a := [3]int{41, 42, 43}
	// 	fmt.Printf("%#v\n", a)

	// 	b := [...]string{"hello", "my", "friends"}
	// 	fmt.Printf("%v", b)

	// 	var c [2]int
	// 	c[0] = 1
	// 	c[1] = 2
	// 	fmt.Println(c)

	// 	fmt.Printf("a is of type %T\tbut the length is %v\n", a, len(a))
	// 	fmt.Printf("b is of type %T\tbut the length is %v\n", b, len(b))
	// 	fmt.Printf("c is of type %T\tbut the length is %v\n", c, len(c))

	// 	myArray := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	// 	fmt.Printf("%T", myArray)
	// }

	{
		xi := [8]int{10, 20, 30, 40}
		for i, v := range xi {
			fmt.Printf("at index %v: %v\n", i, v)
		}
		fmt.Printf("%v\n", len(xi[2:]))
		// xi = append(xi, 50, 60)
		fmt.Printf("%v\n", cap(xi[2:]))

	}

	// delete from slice
	// {
	// 	fmt.Println("before change-------------")
	// 	xi:=[]int{00,10,20,30,40,50,60,70}
	// 	fmt.Println(xi)
	// 	fmt.Printf("the content of xi2 is: %v\n",xi)
	// 	fmt.Printf("the lenght is: %v\n",len(xi))
	// 	fmt.Printf("the initial capacity is: %v\n",cap(xi))
	// 	// remove 40
	// 	fmt.Println("after change-------------")
	// 	// xi=append(xi[:4], xi[5:]...)
	// 	xi = append(xi, 80,90)
	// 	fmt.Println(xi)
	// 	fmt.Printf("the content of xi2 is: %v\n",xi)
	// 	fmt.Printf("the lenght is: %v\n",len(xi))
	// 	fmt.Printf("the initial capacity is: %v\n",cap(xi))
	// 	fmt.Println("xi2-------------")
	// 	xi2:=make([]int,2,10)
	// 	fmt.Printf("the content of xi2 is: %v\n",xi2)
	// 	fmt.Printf("the lenght is: %v\n",len(xi2))
	// 	fmt.Printf("the initial capacity is: %v\n",cap(xi2))
	// }

	{
		xc:= []string{"black","white","red","green"}
		xf:=[]string{"mango","fufu","pizza","balango"}

		xxw:=[][]string{xc,xf}
		fmt.Println(xxw)
	}

}
