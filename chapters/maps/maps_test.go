package maps_test

import (
	"learn-go-with-tests/chapters/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := maps.Dictionary{"test": "this is just a test"}
	t.Run("word not in dictionary", func(t *testing.T) {

		got, err := dictionary.Search("test2")
		assertError(t, err, maps.ErrWordNotFound)
		assertEqualStrings(t, got, "")
	})
	t.Run("word is in dictionary", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertEqualStrings(t, got, want)
		assertNoError(t, err)
	})

}

func TestAdd(t *testing.T) {
	dictionary := maps.Dictionary{}

	dictionary.Add("test3", "this is just a test")
	want := "this is just a test"

	got, err := dictionary.Search("test3")

	assertEqualStrings(t, got, want)
	assertNoError(t, err)
}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but should not have one")
	}
}
