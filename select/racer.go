package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	startUrl1 := time.Now()
	http.Get(url1)
	url1Duration := time.Since(startUrl1)

	startUrl2 := time.Now()
	http.Get(url2)
	url2Duration := time.Since(startUrl2)

	if url1Duration < url2Duration {
		return url1
	}
	return url2
}

func main() {
	fmt.Printf("winner: %s", Racer("http://www.facebook.com", "http://www.quii.co.uk"))
}
