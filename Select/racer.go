package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	winner, _ := Racer("http://www.facebook.com", "http://www.quii.dev")
	fmt.Println(winner)
}

// func Racer(url1, url2 string) (winner string) {
// 	aDuration := measureResponseTime(url1)
// 	bDuration := measureResponseTime(url2)

// 	if aDuration < bDuration {
// 		return url1
// 	}

// 	return url2
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
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
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
