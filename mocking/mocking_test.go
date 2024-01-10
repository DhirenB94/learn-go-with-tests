package main

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints the expected output of 3 2 1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
	2
	1
	Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("sleeps are called in the correct order", func(t *testing.T) {
		spySleeper := &SpyCountdownOperations{}
		Countdown(spySleeper, spySleeper)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
		}
	})
}
