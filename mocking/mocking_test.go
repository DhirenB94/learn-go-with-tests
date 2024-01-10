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

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("sleep is called the correct number of times", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{Calls: 0}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
	2
	1
	Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to the sleeper, want 3 but got %d", spySleeper.Calls)
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
