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
	t.Run("new word", func(t *testing.T) {
		dictionary := maps.Dictionary{}

		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("word already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := maps.Dictionary{word: definition}

		err := dictionary.Add(word, "new definition for existing word")

		assertError(t, err, maps.ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exists, so update", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition of existing word"

		dictionary := maps.Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word, so do not update", func(t *testing.T) {
		word := "test"
		newDefinition := "this is just a test"

		dictionary := maps.Dictionary{}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, maps.ErrNoWordFoundToUpdate)

	})
}
func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := maps.Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, maps.ErrWordNotFound)
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
func assertDefinition(t testing.TB, dictionary maps.Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertNoError(t, err)
	assertEqualStrings(t, got, definition)
}
