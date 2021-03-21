package main

import (
	"fmt"
	"net/http"
	"time"
)

// Racer takes 2 URLs and sends both HTTP GET requests, returning the URL which returned first.
func Racer(url1, url2 string) (winner string) {
	if responseTime(url1) < responseTime(url2) {
		return url1
	}
	return url2
}

func responseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	timeTaken := time.Since(start)
	return timeTaken
}

func main() {
	fmt.Printf("winner: %s", Racer("http://www.facebook.com", "http://www.quii.co.uk"))
}
