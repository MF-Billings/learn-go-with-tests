package concurrency

type WebsiteChecker func(string) bool

// CheckWebsites checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value - true for a good response and false for a bad one.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// anonymous functions maintain access to the lexical scope they're defined in - all the variables
		// that are available at the point when you declare the anonymous function are also available in
		// the body of the function
		go func() {
			results[url] = wc(url)
		}()
	}

	return results
}
