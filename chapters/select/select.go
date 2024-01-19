package selct

import (
	"fmt"
	"net/http"
	"time"
)

//Make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first.
//If none of them return within 10 seconds then it should return an error.

// instead of testing speeds of the sites one after the other, check both at the same time with channels
// exact timings dont matter, just care about which comes first
func WebsiteRacer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

//SELECT AND CHANNELS
// variable := <- channel ----> this is a BLOCKING call as you are wating for a value
// select allows us to wait on multiple channels, the first one to recieve a value "wins" and the code under the "case" is executed - in this case return
// we are setting up 2 channels via the ping func, which ever channel finishes the http.get call 1st, unblocks and executes.

// ping creates a channel struct and returns it when it recieves a signal (closing of the channel)
// it is a channel struct because we dont care about the type thats sent to the channel
// and channel struct is the smallest data type from a memory point of view
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	//as soon as http.Get is completed, the channel is instructed to close and returned
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
