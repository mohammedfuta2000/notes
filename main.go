package main

import "fmt"

func main()  {
	// arrays
	{ 
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

	{
		myArray:= [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
		fmt.Printf("%T",myArray)
	}

	
}