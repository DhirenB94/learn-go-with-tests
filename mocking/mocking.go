package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const coutdownStart = 3

func Countdown(writer io.Writer) {
	for i := coutdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(time.Second * 1)
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
