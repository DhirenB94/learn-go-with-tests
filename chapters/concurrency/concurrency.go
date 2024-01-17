package concurrency

import (
	"fmt"
	"time"
)

type WebsiteChecker func(string) bool

// CheckWebsites checks the status of a list of URLs, returning a map of each url checked to a boolean value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		fmt.Println("out the go routine, url: ", url)
		//each iteration of the loop will run in its own go routine
		go func() {
			fmt.Println("in the go routine, url: ", url)
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}

//Anonymous functions
//The () at the end of the anon func allows them to be executed at the same time they are declared (as opposed to habving to call them somewhere)
// All the variables that are available at the point when you declare the anonymoous function, are also available in the body of the func
