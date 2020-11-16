package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites returns map of each url to boolean value
// true for good response, false for bad response
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
