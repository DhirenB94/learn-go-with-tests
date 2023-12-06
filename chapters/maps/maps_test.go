package maps_test

import (
	"learn-go-with-tests/chapters/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := maps.Search(dictionary, "test")
	want := "this is just a test"

	assertEqualStrings(t, got, want)
}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
