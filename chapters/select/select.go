package selct

import (
	"net/http"
	"time"
)

//Make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first.
//If none of them return within 10 seconds then it should return an error.

func WebsiteRacer(a, b string) string {
	aDuration := measureDuration(a)
	bDuration := measureDuration(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureDuration(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	duration := time.Since(startTime)

	return duration
}
