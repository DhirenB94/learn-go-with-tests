package dep

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	//PrintF writes to stdout, but we need to write to the buffer
	//fmt.Printf("Hello, %s", name)

	//Fprintf allows you to specify the writer you want to write to
	fmt.Fprintf(writer, "Hello, %s", name)
}
