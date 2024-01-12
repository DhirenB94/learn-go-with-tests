package dep_test

import (
	"bytes"
	dep "learn-go-with-tests/dependency-injection"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	dep.Greet(&buffer, "Dhiren")

	got := buffer.String()
	want := "Hello, Dhiren"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
