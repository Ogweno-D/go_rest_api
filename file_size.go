package main

import (
	"fmt"
	"io"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.microsoft.com",
	}
	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Println(err)
				}
			}(res.Body)
			bs, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}
	var biggest HomePageSize
	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Printf("The biggest home page: %s, : %d\n", biggest.URL, biggest.Size)
}
