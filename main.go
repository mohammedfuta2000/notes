package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	first string
}

// mimicing go internals - like logInfo in v1.0.6
func (p Person) writeOut(w io.Writer) {
	 w.Write([]byte(p.first))

}

func main() {
	p:=Person{first: "Ginny"}
	f,err:=os.Create("output.txt");if err!=nil{log.Fatal(err)};defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())

	// how is the buffer memory space released
}
