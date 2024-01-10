package dep

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	//PrintF writes to stdout, but we need to write to the buffer
	//fmt.Printf("Hello, %s", name)

	//Fprintf allows you to specify the writer you want to write to
	fmt.Fprintf(writer, "Hello, %s", name)
}
