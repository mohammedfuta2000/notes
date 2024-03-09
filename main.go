package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// type set
type MyTS interface{
	// int | float64
	// ~int | ~float64
	constraints.Integer | constraints.Float
}

// generics
func add[T MyTS,U comparable](a, b T,c U) T {
	return a+b
}

// type alias and underlying type constraints
type seconds int

type minutes = int

/*
type MyStruct struct {
	// Frequently accessed fields first
	ID   int
	Name string

	// Group related fields together
	Location struct {
		Latitude  float64
		Longitude float64
	}

	// Less frequently accessed or larger fields
	Data []byte
	Flag bool
}

func main() {
	m := MyStruct{
		ID:   1,
		Name: "futa",
		Location: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  12.22,
			Longitude: 12.22},
		Data: []byte("gold"),
		Flag: true}
	m.Location.Latitude=12.33
}
*/

func main()  {
	var n seconds = 20

	// man:=map[MyTS]string{}

	// var m minutes = 10
	fmt.Println(add(n,2,"goof"))
}
