package selct_test

import (
	selct "learn-go-with-tests/chapters/select"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	t.Run("compares the speed of the server, returning the url of the faster one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		fastestURL, err := selct.WebsiteRacer(slowURL, fastURL)

		if fastestURL != fastURL {
			t.Errorf("got %q but want %q", fastestURL, fastURL)
		}

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
	})
	t.Run("returns an error if the server doesnt respond within 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(12 * time.Second)
		serverB := makeDelayedServer(11 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := selct.WebsiteRacer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
