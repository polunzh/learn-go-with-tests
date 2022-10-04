package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	ans := make(map[string]bool, len(urls))
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, checker(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <- resultChannel
		ans[r.string] = r.bool
	}

	return ans
}
