package concurrency_test

import (
	"learn-go-with-tests/chapters/concurrency"
	"reflect"
	"testing"
)

// using dependency injection to test the func without making any http calls
func mockWebsiteChecker(url string) bool {
	if url == "www.abcd.gef" {
		return false
	}
	return true
}

func Test(t *testing.T) {
	urls := []string{
		"www.google.com",
		"www.youtube.com",
		"www.abcd.gef",
	}

	got := concurrency.CheckWebsites(mockWebsiteChecker, urls)

	want := map[string]bool{
		"www.google.com":  true,
		"www.youtube.com": true,
		"www.abcd.gef":    false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but wanted %v", got, want)
	}
}
