package maps_test

import (
	"learn-go-with-tests/chapters/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := maps.Dictionary{"test": "this is just a test"}
	t.Run("word not in dictionary", func(t *testing.T) {

		_, err := dictionary.Search("test2")
		if err == nil {
			t.Fatal("expected an error but did not get one")
		}
		want := "word not in dictionary"

		assertEqualStrings(t, err.Error(), want)
		// Errors can be converted to a string with the .Error() method
	})
	t.Run("word is in dictionary", func(t *testing.T) {
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should not have an error but got one")
		}
		want := "this is just a test"

		assertEqualStrings(t, got, want)
	})

}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
