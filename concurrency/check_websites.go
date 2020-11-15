package concurrency

type WebsiteChecker func(string) bool

// CheckWebsites returns map of each url to boolean value
// true for good response, false for bad response
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
