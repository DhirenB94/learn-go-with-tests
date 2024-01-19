package selct_test

import (
	selct "learn-go-with-tests/chapters/select"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	fastestURL := selct.WebsiteRacer(slowURL, fastURL)

	if fastestURL != fastURL {
		t.Errorf("got %q but want %q", fastestURL, fastURL)
	}

	slowServer.Close()
	fastServer.Close()
}

//using madeup/real urls in the test means that you actually have to hit those and so you are at the mercy of their timings
//we can use the httptest pkg to create a mock HTTP server to test against
//it finds an open port to listen on and then you can close it when you're done with your test.
