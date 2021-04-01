package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer takes 2 URLs and sends both HTTP GET requests, returning the URL which returned first.
func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// don't care why type is sent to the channel, just that there's a signal when it's done and closing the channel fits the bill
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func main() {
	winner, _ := Racer("http://www.facebook.com", "http://www.quii.co.uk")
	fmt.Printf("winner: %s", winner)
}
