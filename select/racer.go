package main

import (
	"fmt"
	"net/http"
)

// Racer takes 2 URLs and sends both HTTP GET requests, returning the URL which returned first.
func Racer(url1, url2 string) (winner string) {
	// select lets you wait on multiple channels.  The first one to send a value "wins" and the code underneath thhe case is executed.
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
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
	fmt.Printf("winner: %s", Racer("http://www.facebook.com", "http://www.quii.co.uk"))
}
