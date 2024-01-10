package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const coutdownStart = 3

func Countdown(writer io.Writer) {
	for i := coutdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
