package selct

import (
	"fmt"
	"net/http"
	"time"
)

//Make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first.
//If none of them return within 10 seconds then it should return an error.

func WebsiteRacer(a, b string) string {
	aStartTime := time.Now()
	//this returns a http.response and an error, but not interested in these for the moment
	http.Get(a)
	aDuration := time.Since(aStartTime)

	bStartTime := time.Now()
	http.Get(b)
	bDuration := time.Since(bStartTime)

	fmt.Println("a", aDuration, "b", bDuration)

	if aDuration < bDuration {
		return a
	}
	return b
}
