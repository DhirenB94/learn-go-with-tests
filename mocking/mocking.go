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

// Allows us to configure the time slept
// duration to configure the time slept and sleep as a way to pass in a sleep function.
// The signature of sleep is the same as for time.Sleep allowing us to use time.Sleep in our real implementation
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
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
