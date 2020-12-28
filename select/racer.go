package main

import "fmt"
import "net/http"
import "time"

// func Racer(a, b string) (winner string) {
// 	// startA := time.Now()
// 	// http.Get(a)
// 	aDuration := measureResponseTime(a)

// 	// startB := time.Now()
// 	// http.Get(b)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
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

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func main() {
	fmt.Println("x")
}
