package main

import (
	"fmt"
	"strconv"
)

type book struct{
	title string
}

func (b book) String() string{
	return fmt.Sprintf("The title of the Book is %v",b.title)
}

type count int

func (c count) String() string{
	return fmt.Sprintf("This is the number %v",strconv.Itoa(int(c)) )
}

// now book and count are also of type stringer, so go itnernal will print them as depicted here



func main()  {
	b:= book{
		title: "gold and water",
	}
	var c count = 4

	fmt.Printf("%v\n",b)
	x:= fmt.Sprintln(c)
	fmt.Println(x)
	// notice how differntlyy it println prints them

}