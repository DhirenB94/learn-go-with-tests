package concurrency

type WebsiteChecker func(string) bool

// CheckWebsites checks the status of a list of URLs, returning a map of each url checked to a boolean value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
