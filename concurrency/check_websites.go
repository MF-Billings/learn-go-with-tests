package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value - true for a good response and false for a bad one.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// anonymous functions maintain access to the lexical scope they're defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
