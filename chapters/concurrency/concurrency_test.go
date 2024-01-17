package concurrency_test

import (
	"learn-go-with-tests/chapters/concurrency"
	"reflect"
	"testing"
	"time"
)

// using dependency injection to test the func without making any http calls
func mockWebsiteChecker(url string) bool {
	if url == "www.abcd.gef" {
		return false
	}
	return true
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// the benchmark tests checkWebsites with 100 urls and a fake websiteChecker (which is deliberately slow)
// resetTimer is used to reset the timing within the test before it actually runs
// takes approx 2.25s

// now with channnels it takes approx 0.025s
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[0] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
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
