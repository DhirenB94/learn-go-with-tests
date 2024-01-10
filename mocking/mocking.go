package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Defining the sleep dependency as an interface
// lets us use the real sleep in main and a spy sleeper in the tests
type Sleeper interface {
	Sleep()
}

type realSleeper struct{}

func (rs *realSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const coutdownStart = 3

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := coutdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &realSleeper{}
	Countdown(os.Stdout, sleeper)
}
