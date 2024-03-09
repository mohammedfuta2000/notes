package main

import "fmt"

type dog struct{
	name string
}

func (d dog)walk()  {
	fmt.Println("my name is ",d.name," and i like to walk")
}
func (d *dog)run()  {
	n2:=d.name
	d.name="rudolf"
	fmt.Printf("i am no longer %v, now i am %v and i love to run!!!\n",n2,d.name)
}

type youngin interface{
	walk()
	run()
}

func youngRun(y youngin) {
	fmt.Println("young young young")
	y.run()
}


func main()  {
	d1:=dog{name: "Henry"}
	d1.walk()
	d1.run()
	// this is NOT possible
	// youngRun(d1)
	

	d2:=&dog{name: "Padget"}
	d2.walk()
	d2.run()
	// this is possible
	youngRun(d2)


}