package selct_test

import (
	selct "learn-go-with-tests/chapters/select"
	"testing"
)

func TestWebsiteRacer(t *testing.T) {
	slowURL := "www.slowurl.com"
	fastURL := "www.fasturl.com"

	fastestURL := selct.WebsiteRacer(slowURL, fastURL)

	if fastestURL != fastURL {
		t.Errorf("got %q but want %q", fastestURL, fastURL)
	}
}
