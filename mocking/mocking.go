package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	fmt.Fprint(writer, "3")
}

func main () {
	Countdown(os.Stdout)
}