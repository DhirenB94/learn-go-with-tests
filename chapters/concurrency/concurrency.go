package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLs, returning a map of each url checked to a boolean value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	resultsMap := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		//each iteration of the loop will run in its own go routine
		//instead of writing directly to the map, we send a "results" struct to the "resultsChannel"
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	//recieve from the channel len(url) times then add to the map
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		resultsMap[r.string] = r.bool
	}

	return resultsMap
}

//Anonymous functions
//The () at the end of the anon func allows them to be executed at the same time they are declared (as opposed to habving to call them somewhere)
// All the variables that are available at the point when you declare the anonymoous function, are also available in the body of the func

// By giving each anonymous function a parameter for the url - u -
// and then calling the anonymous function with the url as the argument
// we make sure that the value of u is fixed as the value of url for the iteration of the loop that we're launching the goroutine in.
// u is a copy of the value of url, and so can't be changed.

//RACE CONDITIONS
// With the above code, sometimes 2 of the go routines will try to write to the map at the exact same time (see with benchmark tests)
// This is a race condition - a bug that occurs when the output of our software is dependent on timing and sequence of events that we cannot control
// can detect race conditions with go test -race

//CHANNELS
//we can solve the above issues by coordinating our go routines using channles
//these are a Go Data Structure that can send and recieve values, to allow communication between differet processes

//Send Statement
// channel <- value

//Recieve Statement
//variable := <- channel
