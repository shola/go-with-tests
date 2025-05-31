package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	// `select` allows you to wait on multiple channels, and the
	// first one wins
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// an empty struct uses the least amount of memory of any
// datatype, it's used because it's the most lightweight
func ping(url string) chan struct{} {
	// channels must be created with make. create a channel as
	// a var without init will make it nil, which will block
	// forever
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}