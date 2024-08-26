package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, map_url2success *map[string]bool, wg *sync.WaitGroup, mu *sync.Mutex) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	_, exists := (*map_url2success)[url]
	mu.Unlock()

	if !exists {
		mu.Lock()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			(*map_url2success)[url] = false
			mu.Unlock()
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		(*map_url2success)[url] = true
		mu.Unlock()
		for _, u := range urls {
			Crawl(u, depth-1, fetcher, map_url2success, wg, mu)
		}
	}
	return
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	map_url := make(map[string]bool)
	Crawl("https://golang.org/", 4, fetcher, &map_url, &wg, &mu)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
